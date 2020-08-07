/*
@Time : 2020/7/15 11:39 下午
@Author : lucbine
@File : user.go
*/
package service

import "github.com/gin-gonic/gin"

type UserService struct {
	ctx *gin.Context
}

func NewUserService(c *gin.Context) *UserService {
	return &UserService{
		ctx: c,
	}
}

//
