/*
@Time : 2020/7/13 4:35 下午
@Author : lucbine
@File : cron.go
*/
package service

import (
	"hotwheels/agent/constant"
	"hotwheels/agent/internal/config"
	"hotwheels/agent/internal/logger"
	"sync"
	"time"

	"go.uber.org/zap"

	"errors"

	"github.com/robfig/cron/v3"
)

var (
	jobCron *cron.Cron
	lock    sync.Mutex
	//workPool =
	//定义任务列表
	JobList = make([]*Job, 0)
)

func init() {
	//初始化定位任务
	jobCron = cron.New()
	jobCron.Start()
}

//添加定时任务
func AddJob(spec string, job *Job) (bool, error) {
	lock.Lock()
	defer lock.Unlock()
	if GetEntryById(job.Id) != nil {
		logger.Logger.Warn("job is already exists", zap.Int(constant.JobId, job.Id))
		return false, errors.New("job is already exists")
	}

	if _, err := jobCron.AddJob(spec, job); err != nil {
		logger.Logger.Error("job is add fail", zap.Int(constant.JobId, job.Id), zap.Any("job", job), zap.Error(err))
		return false, err
	}
	logger.Logger.Info("job is add success", zap.Int(constant.JobId, job.Id))
	return true, nil
}

//删除定时任务
func DelJob(id cron.EntryID) bool {
	lock.Lock()
	defer lock.Unlock()
	jobCron.Remove(id)
	//todo 支持强杀逻辑
	return jobCron.Entry(id).Valid() == false
}

//获得实体
func GetEntryById(id int) *cron.Entry {
	for _, e := range jobCron.Entries() {
		if v, ok := e.Job.(*Job); ok {
			if v.Id == id {
				return &e
			}
		}
	}
	return nil
}

//定期拉取任务
func InitCronJob() {
	client := NewServerClient()
	c := cron.New(cron.WithSeconds())
	c.AddFunc(config.GetString("config.cron.checkInterval"), func() {
		jobList, err := client.JobList()
		if err != nil {
			logger.Logger.Error(constant.JobGetList, zap.Error(err))
			return
		}

		if len(jobList) == 0 {
			logger.Logger.Warn(constant.JobGetListEmpty)
			return
		}
		//加入任务
		for _, item := range jobList {
			AddJob(item.Spec, &Job{
				Id:              item.Id,
				Name:            item.Name,
				Spec:            item.Spec,
				Command:         item.Command,
				Timeout:         time.Duration(item.Timeout) * time.Second,
				Concurrent:      item.Concurrent,
				ConcurrentCount: item.ConcurrentCount,
			})
		}
		return
	})
	c.Start()
}
