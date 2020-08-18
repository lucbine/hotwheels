/*
@Time : 2020/7/13 3:17 下午
@Author : lucbine
@File : jobs.go
*/
package service

import (
	"bytes"
	"fmt"
	"hotwheels/agent/constant"
	"hotwheels/agent/entity"
	"hotwheels/agent/internal/core"
	"hotwheels/agent/internal/logger"
	"os"
	"os/exec"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

//https://www.cnblogs.com/jiangz222/p/12345566.html

const (
	jobCountLua = `if redis.call("get", KEYS[1]) < ARGV[1] then return redis.call("incr", KEYS[1]) else return 0 end`
)

//定时去服务中心拉任务
type Job struct { //任务
	Id              int                                                                                    //任务id
	Name            string                                                                                 //任务名称
	User            string                                                                                 //指定运行用户
	Spec            string                                                                                 //定时规则
	Command         string                                                                                 //脚本执行命令
	Status          int                                                                                    //任务状态，大于0 表示正在执行中
	Timeout         time.Duration                                                                          //超时
	RunFunc         func(duration time.Duration) (bufOut string, bufErr string, err error, isTimeout bool) //任务调度函数
	Concurrent      bool                                                                                   //同一台机器是否允许并行执行                                                                             //同一个任务是否允许并行执行
	ConcurrentCount int                                                                                    //并行执行的个数
	ForceKill       bool                                                                                   //是否强制杀掉脚本进程
	Process         *os.Process                                                                            //任务运行进程
}

func NewJob(task *entity.Task) *Job {
	j := &Job{
		Id:              task.Id,
		Name:            task.Name,
		User:            task.User,
		Spec:            task.Spec,
		Timeout:         time.Duration(task.Timeout) * time.Second,
		Command:         task.Command,
		ForceKill:       task.ForceKill,
		Concurrent:      task.Concurrent,
		ConcurrentCount: task.ConcurrentCount,
	}
	j.RunFunc = func(timeout time.Duration) (bufOut string, bufErr string, err error, isTimeout bool) {
		var (
			outBuf = new(bytes.Buffer)
			errBuf = new(bytes.Buffer)
		)
		cmd := exec.Command("/bin/bash", "-c", j.Command)
		cmd.Stdout = outBuf
		cmd.Stderr = errBuf

		//fmt.Println("j.user", j.User)
		//指定用户运行
		//if j.User != "" {
		//	user, err := user.Lookup(j.User)
		//
		//	fmt.Printf("uid=%s,gid=%s \n", user.Uid, user.Gid)
		//	if err != nil {
		//		return bufOut, bufErr, errors.New(fmt.Sprintf("没有找到此用户：%s", j.User)), isTimeout
		//	}
		//	uid, _ := strconv.Atoi(user.Uid)
		//	gid, _ := strconv.Atoi(user.Gid)
		//	cmd.SysProcAttr = &syscall.SysProcAttr{}
		//	cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)}
		//}
		cmd.Start()
		j.Process = cmd.Process
		err, isTimeout = runCmdWithTimeout(j.Id, cmd, timeout)
		return outBuf.String(), errBuf.String(), err, isTimeout
	}
	return j
}

func runCmdWithTimeout(jobId int, cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		logger.Logger.Warn(constant.JobTimeout, zap.Int(constant.JobId, jobId), zap.Int(constant.JobProcessId, cmd.Process.Pid),
			zap.Int(constant.JobTimeoutSecond, int(timeout/time.Second)))
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			logger.Logger.Error(constant.JobNotStop, zap.Int(constant.JobId, jobId), zap.Int(constant.JobProcessId, cmd.Process.Pid), zap.Error(err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}

//实现job 执行方法
func (j *Job) Run() {
	//过期时间
	t := time.Now()
	timeout := time.Duration(24 * time.Hour) //如果没有设置超时时间 默认1天
	if j.Timeout.Seconds() > 0 {
		timeout = j.Timeout
	}

	localKey := fmt.Sprintf(constant.JobLock, j.Id)
	//先获得一个锁

	ok, err := core.LocalCache.Eval(jobCountLua, []string{localKey}, 5).Int()
	//ok, err := core.LocalCache.SetNX(localKey, 1, timeout).Result()
	if ok == 0 || err != nil {
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	var execResult = new(entity.ExecResult)
	execResult.TaskId = j.Id
	//执行完  进行日志上报
	defer Report(execResult)
	defer core.LocalCache.Incr(localKey) //减少一个 计数
	//异常捕获
	defer func() {
		if err := recover(); err != nil {
			logger.Logger.Error(constant.JobRunErr, zap.Int(constant.JobId, j.Id), zap.String("err stack", string(debug.Stack())))
		}
	}()
	if j.Status == constant.JobStatusRunable && !j.Concurrent {
		logger.Logger.Warn(constant.JobRunning, zap.Int(constant.JobId, j.Id))
		return
	}
	if j.Status == constant.JobStatusRunable && j.Concurrent && j.ConcurrentCount > 10 { //todo
		logger.Logger.Warn(constant.JobRunning, zap.Int(constant.JobId, j.Id))
		return
	}

	cmdOut, cmdErr, err, isTimeout := j.RunFunc(timeout)

	execTime := time.Now().Sub(t) / time.Millisecond

	execResult.Output = cmdOut
	execResult.Error = cmdErr
	execResult.ProcessTime = int64(execTime)
	execResult.CreateTime = time.Now().Unix()
	execResult.Status = constant.JobStatusSuccess

	if isTimeout {
		execResult.Status = constant.JobStatusTimeout
		execResult.Error = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), cmdErr)
	} else if err != nil {
		execResult.Status = constant.JobStatusError
		execResult.Error = err.Error() + ":" + cmdErr
	}
	return
}

//运行日志上报
func Report(execResult *entity.ExecResult) {
	go func() {
		if err := recover(); err != nil {
			logger.Logger.Error(constant.JobRunErr, zap.Int(constant.JobId, execResult.TaskId), zap.String("err stack", string(debug.Stack())))
		}
		//上报
		//上报结果
		err := NewServerClient().Result(execResult)
		if err != nil {
			logger.Logger.Error(constant.JobResultReport, zap.Int(constant.JobId, execResult.TaskId), zap.Any("exec_result", execResult), zap.Error(err))
			return
		}
		logger.Logger.Info(constant.JobResultReport, zap.Int(constant.JobId, execResult.TaskId), zap.Any("exec_result", execResult))
	}()
}
