/*
@Time : 2020/7/16 11:20 上午
@Author : lucbine
@File : index.go
*/
package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//任务列表
func GroupList(ctx *gin.Context) {
	templateParams := BaseC(ctx, "任务列表")
	ctx.HTML(http.StatusOK, "group/list.html", templateParams)
}

//新增任务
func GroupAdd(ctx *gin.Context) {
	templateParams := BaseC(ctx, "任务列表")
	ctx.HTML(http.StatusOK, "group/add.html", templateParams)
}

//
func GroupEdit(ctx *gin.Context) {
	templateParams := BaseC(ctx, "任务列表")
	ctx.HTML(http.StatusOK, "group/add.html", templateParams)
}
