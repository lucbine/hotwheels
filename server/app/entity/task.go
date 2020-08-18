/*
@Time : 2020/7/15 10:44 下午
@Author : lucbine
@File : job.go
*/
package entity

type SetHotContentReq struct {
	ContentId      int    `json:"content_id" form:"content_id"`
	TopTopicIds    int    `json:"top_topic_ids" form:"top_topic_ids"`
	TopStartTime   int    `json:"top_start_time" form:"top_start_time"`
	TopEndTime     int    `json:"top_end_time" form:"top_end_time"`
	RecTopicIds    string `json:"rec_topic_ids" form:"rec_topic_ids"`
	RecTabTopicIds string `json:"rec_tab_topic_ids" form:"rec_tab_topic_ids"`
	OpId           int    `json:"op_id" form:"op_id"`
	OpName         string `json:"op_name" form:"op_name"`
}

type AddTaskReq struct {
	GroupId         int    `form:"group_id" json:"group_id"`
	TaskName        string `form:"task_name" json:"task_name"`
	TaskType        string `form:"task_type" json:"task_type"`
	Description     string `form:"description" json:"description"`
	CronSpec        string `form:"cron_spec" json:"cron_spec"`
	Concurrent      string `form:"concurrent" json:"concurrent"`
	ConcurrentCount string `form:"concurrent_count" json:"concurrent_count"`
	Command         string `form:"command" json:"command"`
}

type ReportReq struct {
	TaskId      int64  `form:"task_id" json:"task_id"`
	Output      string `form:"output" json:"output"`
	Error       string `form:"error" json:"error"`
	Status      int8   `form:"status" json:"status"`
	ProcessTime int    `form:"process_time" json:"process_time"`
}

type HcReq struct {
	IP          string  `form:"ip" json:"ip" binding:"required"`
	HostName    string  `form:"host_name" json:"host_name" binding:"required"`
	CpuCount    int     `form:"cpu_count" json:"cpu_count" binding:"required"`
	Os          string  `form:"os" json:"os" binding:"required"`
	CpuUsage    float64 `form:"cpu_usage" json:"cpu_usage" binding:"required"`
	MemorySize  uint64  `form:"memory_size" json:"memory_size" binding:"required"`
	MemoryUsage float64 `form:"memory_usage" json:"memory_usage" binding:"required"`
}
