/*
@Time : 2020/7/13 4:35 下午
@Author : lucbine
@File : cron.go
*/
package service

import (
	"hotwheels/agent/constant"
	"hotwheels/agent/internal/logger"
	"sync"

	"go.uber.org/zap"

	"github.com/robfig/cron/v3"
)

var (
	jobCron *cron.Cron
	lock    sync.Mutex
	JobList = make([]*Job, 0)
)

func init() {
	//初始化定位任务
	jobCron = cron.New(cron.WithSeconds())
	jobCron.Start()
}

//添加定时任务
func AddJob(spec string, job *Job) (bool, error) {
	lock.Lock()
	defer lock.Unlock()
	entryId, entryJob := GetJobById(job.Id)
	if entryJob != nil {
		//强杀
		if job.ForceKill {
			return ForceDelJob(entryId, entryJob), nil
		}

		//删除定时任务
		if job.Status == constant.JobStatusStop {
			return DelJob(entryId), nil
		}

		//比较有没有变化
		if job.Spec != entryJob.Spec || job.Command != entryJob.Command || job.Timeout != entryJob.Timeout {
			DelJob(entryId) //有变化  则先从entry 里删除   再添加
		}
	}
	if _, err := jobCron.AddJob(spec, job); err != nil {
		logger.Logger.Error("job is add fail", zap.Int(constant.JobId, job.Id), zap.Any("job", job), zap.Error(err))
		return false, err
	}
	logger.Logger.Info("job is add success", zap.Int(constant.JobId, job.Id))
	return true, nil
}

//删除定时任务（不支持强杀）
func DelJob(id cron.EntryID) bool {
	lock.Lock()
	defer lock.Unlock()
	//删除entry
	jobCron.Remove(id)
	//杀死进程
	return jobCron.Entry(id).Valid() == false
}

//强杀
func ForceDelJob(id cron.EntryID, entryJob *Job) bool {
	lock.Lock()
	defer lock.Unlock()
	//删除entry
	jobCron.Remove(id)
	//杀死进程
	err := entryJob.Process.Kill()
	if err != nil {
		logger.Logger.Error("job force kill is fail", zap.Int(constant.JobId, entryJob.Id), zap.Any("runJob", entryJob))
	}
	return jobCron.Entry(id).Valid() == false
}

//获得job
func GetJobById(id int) (cron.EntryID, *Job) {
	for _, e := range jobCron.Entries() {
		if v, ok := e.Job.(*Job); ok {
			if v.Id == id {
				return e.ID, v
			}
		}
	}
	return 0, nil
}

//定期拉取任务
func InitCronJob(cron *cron.Cron) {

	client := NewServerClient()
	taskList, err := client.TaskList()
	if err != nil {
		logger.Logger.Error(constant.JobGetList, zap.Error(err))
		return
	}
	if len(taskList) == 0 {
		logger.Logger.Warn(constant.JobGetListEmpty)
		return
	}

	//加入任务
	for _, item := range taskList {
		AddJob(item.Spec, NewJob(&item))
	}
	return

	//cron.AddFunc(config.GetString("config.cron.checkInterval"), func() {
	//	client := NewServerClient()
	//	taskList, err := client.TaskList()
	//	if err != nil {
	//		logger.Logger.Error(constant.JobGetList, zap.Error(err))
	//		return
	//	}
	//	if len(taskList) == 0 {
	//		logger.Logger.Warn(constant.JobGetListEmpty)
	//		return
	//	}
	//
	//	//加入任务
	//	for _, item := range taskList {
	//		AddJob(item.Spec,  NewJob(&item))
	//	}
	//	return
	//})
	cron.Start()
}
