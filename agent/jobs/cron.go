/*
@Time : 2020/7/13 4:35 下午
@Author : lucbine
@File : cron.go
*/
package jobs

import (
	"sync"

	"errors"

	"github.com/robfig/cron/v3"
)

var (
	jobCron *cron.Cron
	lock    sync.Mutex
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
		return false, errors.New("job is already exists")
	}

	if _, err := jobCron.AddJob(spec, job); err != nil {
		return false, err
	}
	return true, nil
}

//删除定时任务
func DelJob(id cron.EntryID) bool {
	lock.Lock()
	defer lock.Unlock()
	jobCron.Remove(id)
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
