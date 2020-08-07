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
	"hotwheels/agent/internal/logger"
	"os/exec"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

//https://www.cnblogs.com/jiangz222/p/12345566.html
//定时去服务中心拉任务
type Job struct { //任务
	Id              int                                                                                    //任务id
	Name            string                                                                                 //任务名称
	Spec            string                                                                                 //定时规则
	Command         string                                                                                 //脚本执行命令
	Status          int                                                                                    //任务状态，大于0 表示正在执行中
	Timeout         time.Duration                                                                          //超时
	RunFunc         func(duration time.Duration) (bufOut string, bufErr string, err error, isTimeout bool) //任务调度函数
	Concurrent      bool                                                                                   //同一台机器是否允许并行执行                                                                             //同一个任务是否允许并行执行
	ConcurrentCount int                                                                                    //并行执行的个数
}

func NewJob(id int, name string, spec string, command string) *Job {
	j := &Job{
		Id:   id,
		Name: name,
		Spec: spec,
	}
	j.RunFunc = func(timeout time.Duration) (bufOut string, bufErr string, err error, isTimeout bool) {
		var (
			outBuf = new(bytes.Buffer)
			errBuf = new(bytes.Buffer)
		)
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = outBuf
		cmd.Stderr = errBuf

		//fmt.Println(cmd.Process.Pid) //打印进程号
		//cmd.Process.Kill()           //杀死进程

		cmd.Start()
		err, isTimeout = runCmdWithTimeout(id, cmd, timeout)
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
	if j.Status > 0 && !j.Concurrent {
		logger.Logger.Warn(constant.JobRunning, zap.Int(constant.JobId, j.Id))
		return
	}
	if j.Status > 0 && j.Concurrent && j.ConcurrentCount > 10 { //todo
		logger.Logger.Warn(constant.JobRunning, zap.Int(constant.JobId, j.Id))
		return
	}
	defer func() {
		if err := recover(); err != nil {
			logger.Logger.Error(constant.JobRunErr, zap.Int(constant.JobId, j.Id), zap.String("err stack", string(debug.Stack())))
		}
	}()

	t := time.Now()
	timeout := time.Duration(24 * time.Hour)
	if j.Timeout.Seconds() > 0 {
		timeout = j.Timeout
	}
	cmdOut, cmdErr, err, isTimeout := j.RunFunc(timeout)

	execTime := time.Now().Sub(t) / time.Millisecond

	var execResult = &entity.ExecResult{
		JobId:       j.Id,
		StrOut:      cmdOut,
		StrErr:      cmdErr,
		ProcessTime: int64(execTime),
		CreateTime:  time.Now().Unix(),
		Status:      constant.JobStatusSuccess,
	}

	if isTimeout {
		execResult.Status = constant.JobStatusTimeout
		execResult.StrErr = fmt.Sprintf("任务执行超过 %d 秒\n----------------------\n%s\n", int(timeout/time.Second), cmdErr)
	} else if err != nil {
		execResult.Status = constant.JobStatusError
		execResult.StrErr = err.Error() + ":" + cmdErr
	}

	//上报结果
	err = NewServerClient().Result(execResult)
	if err != nil {
		logger.Logger.Error(constant.JobResultReport, zap.Int(constant.JobId, j.Id), zap.Any("exec_result", execResult), zap.Error(err))
		return
	}
	logger.Logger.Info(constant.JobResultReport, zap.Int(constant.JobId, j.Id), zap.Any("exec_result", execResult))
	return
}
