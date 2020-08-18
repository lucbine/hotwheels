/*
@Time : 2020/7/15 11:39 下午
@Author : lucbine
@File : log.go
*/
package service

import (
	"fmt"
	"hotwheels/server/app/entity"
	"hotwheels/server/app/model"
	"hotwheels/server/internal/errcode"

	"github.com/gin-gonic/gin"
)

type LogService struct {
	ctx *gin.Context
}

func NewLogService(c *gin.Context) *LogService {
	return &LogService{
		ctx: c,
	}
}

//上报日志
func (l *LogService) Report(report *entity.ReportReq) *errcode.Err {
	//插入日志
	data := model.HtaskLogModel{
		TaskId:      report.TaskId,
		Output:      report.Output,
		Error:       report.Error,
		Status:      report.Status,
		ProcessTime: report.ProcessTime,
	}

	newId, err := model.NewHtaskLogModel().Add(data)
	if err != nil {
		return errcode.ErrorFail
	}
	//查询告警规则
	fmt.Println(newId)
	//告警通知
	return errcode.Success
}
