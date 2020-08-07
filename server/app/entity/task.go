/*
@Time : 2020/7/15 10:44 下午
@Author : lucbine
@File : job.go
*/
package entity

type AddTaskReq struct {
	GroupId         string `form:"group_id" json:"group_id"`
	TaskName        string `form:"task_name" json:"task_name"`
	TaskType        string `form:"task_type" json:"task_type"`
	Description     string `form:"description" json:"description"`
	CronSpec        string `form:"cron_spec" json:"cron_spec"`
	Concurrent      string `form:"concurrent" json:"concurrent"`
	ConcurrentCount string `form:"concurrent_count" json:"concurrent_count"`
	Command         string `form:"command" json:"command"`
}
