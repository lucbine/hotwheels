/*
@Time : 2020/7/13 4:07 下午
@Author : lucbine
@File : job.go
*/
package constant

//任务相关日志名称规范
const (
	JobId            = "任务ID"
	JobTimeout       = "任务超时"
	JobRunning       = "任务正在执行"
	JobRunErr        = "任务执行失败"
	JobProcessId     = "进程ID"
	JobTimeoutSecond = "超时时间"
	JobNotStop       = "任务无法终止"
	JobResultReport  = "任务执行结果上报"
	JobGetList       = "拉去定时任务失败"
	JobGetListEmpty  = "定时任务为空"
)

const (
	JobStatusSuccess = 0  // 任务执行成功
	JobStatusError   = -1 //任务执行出错
	JobStatusTimeout = -2 //任务执行超时
)

// 时间格式
const (
	TimeFormatDate   = "2006-01-02"              // 日期格式
	TimeFormatDateV2 = "20060102"                // 日期格式V2
	TimeFormatSec    = "2006-01-02 15:04:05"     // 日期秒格式
	TimeFormatMsec   = "2006-01-02 15:04:05.000" // 日期毫秒秒格式
	TimeFormatClock  = "15:04:05"                // 时钟时间格式
)

const (
	JobStatusRunable = 1 //可运行
	JobStatusStop    = 0 //不可用  停止状态

	JobLock = "job_lock_%d"
)
