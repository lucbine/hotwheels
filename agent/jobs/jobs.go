/*
@Time : 2020/7/13 3:17 下午
@Author : lucbine
@File : jobs.go
*/
package jobs

import (
	"bytes"
	"hotwheels/agent/constant"
	"hotwheels/agent/internal/logger"
	"os/exec"
	"runtime/debug"
	"time"

	"go.uber.org/zap"
)

//定时去服务中心拉任务
type Job struct { //任务
	Id              int                                                                                    //任务id
	Name            string                                                                                 //任务名称
	Spec            string                                                                                 //定时规则
	Status          int                                                                                    //任务状态，大于0 表示正在执行中
	RunFunc         func(duration time.Duration) (bufOut string, bufErr string, err error, isTimeout bool) //任务调度函数
	Concurrent      bool                                                                                   //同一台机器是否允许并行执行                                                                             //同一个任务是否允许并行执行
	ConcurrentCount int                                                                                    //并行执行的个数
}

func NewJob(id int, name string, spec string, commend string) *Job {
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
		cmd := exec.Command("/bin/bash", "-c", commend)
		cmd.Stdout = outBuf
		cmd.Stderr = errBuf
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

}
