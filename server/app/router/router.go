/*
@Time : 2020/7/14 9:23 下午
@Author : lucbine
@File : router.go
*/
package router

import (
	"hotwheels/server/app/constant"
	"hotwheels/server/app/controller/api"
	"hotwheels/server/app/controller/web"
	"hotwheels/server/internal/config"
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(runMode())

	//添加健康检查响应,请勿删除
	router.Any("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	apiGroup := router.Group("/api")
	{
		//agent 接口
		apiGroup.GET("/agent/report", api.Report)
		apiGroup.POST("/agent/jobList", api.JobList)

		//admin 接口
		apiGroup.POST("/admin/addTask", api.AddTask)
		apiGroup.POST("/admin/editTask", api.EditTask)
		apiGroup.GET("/admin/taskList", api.TaskList)
		apiGroup.GET("/admin/stat", api.Stat)
	}

	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "view",
		Extension:    ".html",
		Master:       "layout/master",
		DisableCache: true,
	})

	router.GET("/", web.Index)

	return router
}

// 根据env环境变量获取gin运行模式
func runMode() string {
	if config.Env() == constant.EnvPrd {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
