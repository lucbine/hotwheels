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

func BaseC(ctx *gin.Context, pageTitle string) gin.H {
	//初始化页面title
	return gin.H{
		"pageTitle": pageTitle,
		"siteName":  "分布式定时任务管理系统",
		"version":   "v1.2",
	}
}

//首页
func Index(ctx *gin.Context) {
	templateParams := BaseC(ctx, "首页")
	ctx.HTML(http.StatusOK, "main/index.html", templateParams)
}

//登录页
func Login(ctx *gin.Context) {

}

//帮助页
func Help(ctx *gin.Context) {
	templateParams := BaseC(ctx, "帮助")
	ctx.HTML(http.StatusOK, "help/index.html", templateParams)
}
