/*
@Time : 2020/7/15 11:39 下午
@Author : lucbine
@File : group.go
*/
package service

import "github.com/gin-gonic/gin"

type GroupService struct {
	ctx *gin.Context
}

func NewGroupService(c *gin.Context) *GroupService {
	return &GroupService{
		ctx: c,
	}
}

//
