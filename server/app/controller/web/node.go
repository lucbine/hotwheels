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

//节点列表
func NodeList(ctx *gin.Context) {
	templateParams := BaseC(ctx, "节点列表")
	ctx.HTML(http.StatusOK, "node/list.html", templateParams)
}

//节点编辑
func NodeEdit(ctx *gin.Context) {
	templateParams := BaseC(ctx, "编辑节点")
	ctx.HTML(http.StatusOK, "node/edit.html", templateParams)
}
