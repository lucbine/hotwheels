/*
@Time : 2020/7/15 11:39 下午
@Author : lucbine
@File : log.go
*/
package service

import "github.com/gin-gonic/gin"

type LogService struct {
	ctx *gin.Context
}

func NewLogService(c *gin.Context) *LogService {
	return &LogService{
		ctx: c,
	}
}

//
