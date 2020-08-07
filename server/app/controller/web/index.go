/*
@Time : 2020/7/16 11:20 上午
@Author : lucbine
@File : index.go
*/
package web

import (
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{})
}
