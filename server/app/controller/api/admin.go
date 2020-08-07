/*
@Time : 2020/7/14 9:29 下午
@Author : lucbine
@File : admin.go
*/
package api

import (
	"fmt"
	"hotwheels/server/app/entity"
	"hotwheels/server/app/service"
	"hotwheels/server/internal/errcode"
	"hotwheels/server/internal/response"

	"github.com/robfig/cron/v3"

	"github.com/gin-gonic/gin"
)

//添加任务
func AddTask(ctx *gin.Context) {
	var reqs = new(entity.AddTaskReq)
	err := ctx.Bind(reqs)
	if err != nil {
		response.Json(ctx, errcode.InputParamsError, nil)
	}

	reqs.CronSpec = "* * * * * *"
	fmt.Println(reqs)

	//参数校验 todo
	if _, err := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow).Parse(reqs.CronSpec); err != nil {
		response.Json(ctx, errcode.CronFormatError, nil)
		return
	}
	resultErr := service.NewTaskService(ctx).Add(reqs)
	response.Json(ctx, resultErr, nil)
	return
}

//任务编辑
func EditTask(ctx *gin.Context) {

}

//任务列表
func TaskList(ctx *gin.Context) {

}

//统计数据
func Stat(ctx *gin.Context) {

}

//强杀任务
func Kill(ctx *gin.Context) {

}

//暂停任务
func Pause(ctx *gin.Context) {

}
