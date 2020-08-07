/*
@Time : 2020/7/13 9:24 下午
@Author : lucbine
@File : log.go
*/
package entity

type ExecResult struct {
	JobId       int    `json:"job_id"`       //任务id
	StrOut      string `json:"str_out"`      //任务标准输出
	StrErr      string `json:"str_err"`      //任务错误输出
	ProcessTime int64  `json:"process_time"` //执行时间
	CreateTime  int64  `json:"create_time"`  //日志创建时间
	Status      int    `json:"status"`       //执行结果状态
}

type Job struct {
	Id              int    `json:"id"`               //任务id
	Name            string `json:"name"`             //任务名称
	Spec            string `json:"spec"`             //定时规则
	Command         string `json:"command"`          //执行命令
	Timeout         int    `json:"timeout"`          //脚本超时时间
	Concurrent      bool   `json:"concurrent"`       //是否支持并行
	ConcurrentCount int    `json:"concurrent_count"` //支持并行的数量
}
