/*
@Time : 2020/7/13 9:24 下午
@Author : lucbine
@File : log.go
*/
package entity

type ExecResult struct {
	TaskId      int    `json:"task_id"`      //任务id
	Output      string `json:"output"`       //任务标准输出
	Error       string `json:"error"`        //任务错误输出
	ProcessTime int64  `json:"process_time"` //执行时间
	CreateTime  int64  `json:"create_time"`  //日志创建时间
	Status      int    `json:"status"`       //执行结果状态
}

type Task struct {
	Id              int    `json:"id"`               //任务id
	Name            string `json:"name"`             //任务名称
	User            string `json:"user"`             //运行用户
	Spec            string `json:"spec"`             //定时规则
	Command         string `json:"command"`          //执行命令
	Timeout         int    `json:"timeout"`          //脚本超时时间
	Concurrent      bool   `json:"concurrent"`       //是否支持并行
	ConcurrentCount int    `json:"concurrent_count"` //支持并行的数量
	Status          int    `json:"status"`           //任务状态，大于0 表示正在执行中
	ForceKill       bool   `json:"force_kill"`       //是否强制杀掉脚本进程
}

type BaseResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	ShowErr     int    `json:"showErr"`
	CurrentTime int64  `json:"currentTime"`
}
