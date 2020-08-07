/*
@Time : 2020/7/15 10:44 下午
@Author : lucbine
@File : node.go
*/
package service

import (
	"hotwheels/server/internal/errcode"

	"github.com/gin-gonic/gin"
)

type NodeService struct {
	ctx *gin.Context
}

func NewNodeService(c *gin.Context) *NodeService {
	return &NodeService{
		ctx: c,
	}
}

//添加定时任务
func (ts *NodeService) add() *errcode.Err {

	return errcode.Success
}

//定时任务列表
func (ts *NodeService) List() *errcode.Err {

	return errcode.Success
}

//编辑任务
func (ts *NodeService) Edit() *errcode.Err {

	return errcode.Success
}

//任务统计
func (ts *NodeService) Stat() *errcode.Err {

	return errcode.Success
}
