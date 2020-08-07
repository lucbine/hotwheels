/*
@Time : 2020/7/15 10:44 下午
@Author : lucbine
@File : task.go
*/
package service

import (
	"hotwheels/server/app/entity"
	"hotwheels/server/app/model"
	"hotwheels/server/internal/errcode"

	"github.com/gin-gonic/gin"
)

type TaskService struct {
	ctx *gin.Context
}

func NewTaskService(c *gin.Context) *TaskService {
	return &TaskService{
		ctx: c,
	}
}

//添加节点
func (ts *TaskService) Add(data *entity.AddTaskReq) *errcode.Err {
	//逻辑校验 1、 群组校验  2、去重校验

	var insertData = model.HtaskModel{
		GroupId:         data.GroupId,
		TaskName:        data.TaskName,
		TaskType:        data.TaskType,
		Description:     data.Description,
		CronSpec:        data.CronSpec,
		Concurrent:      data.Concurrent,
		ConcurrentCount: data.ConcurrentCount,
		Command:         data.Command,
	}
	model.NewHtaskModel().Add(insertData)
	return errcode.Success
}

//节点列表
func (ts *TaskService) List() *errcode.Err {

	return errcode.Success
}

//节点信息
func (ts *TaskService) Info() *errcode.Err {

	return errcode.Success
}
