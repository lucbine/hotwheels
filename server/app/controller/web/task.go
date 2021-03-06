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
func TaskList(ctx *gin.Context) {
	templateParams := BaseC(ctx, "任务列表")
	ctx.HTML(http.StatusOK, "task/list.html", templateParams)
}

//新增任务
func TaskAdd(ctx *gin.Context) {
	templateParams := BaseC(ctx, "新建任务")
	ctx.HTML(http.StatusOK, "task/add.html", templateParams)
}
