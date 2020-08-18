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

	//api
	apiGroup := router.Group("/api/agent")
	{
		//agent 接口
		apiGroup.POST("/report", api.Report)
		apiGroup.GET("/jobList", api.JobList)
		apiGroup.POST("/hc", api.Hc)
	}

	adminGroup := router.Group("/api/admin")
	{
		//admin 接口
		adminGroup.POST("/addTask", api.AddTask)
		adminGroup.POST("/editTask", api.EditTask)
		adminGroup.GET("/taskList", api.TaskList)
		adminGroup.GET("/stat", api.Stat)
	}
	//web
	//设置静态文件地址
	router.Static("/static", "./static")
	//router.StaticFile("/favicon.ico", "./favicon.ico") //设置图标
	//router.LoadHTMLGlob("layout/*")
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "view",
		Extension: ".html",
		//Master:    "layout/master",
		Partials: []string{
			"layout/header",
			"layout/footer",
		},
		DisableCache: true,
	})

	gin.IsDebugging()

	//首页
	router.GET("/", web.Index)

	//登录页
	router.GET("/login", web.Login)

	//帮助页
	router.GET("/help", web.Help)

	//任务
	taskGroup := router.Group("/task")
	{
		taskGroup.GET("/list", web.TaskList)
		taskGroup.GET("/add", web.TaskAdd)
	}

	//节点
	nodeGroup := router.Group("/node")
	{
		nodeGroup.GET("/list", web.NodeList)
		nodeGroup.GET("/edit", web.NodeEdit)
	}

	//分组
	gGroup := router.Group("/group")
	{
		gGroup.GET("/list", web.GroupList)
		gGroup.GET("/add", web.GroupAdd)
		gGroup.GET("/edit", web.GroupAdd)

	}

	//alarm 告警
	alarmGroup := router.Group("/alarm")
	{
		alarmGroup.GET("/list", web.Login)

	}

	//日志
	logGroup := router.Group("/log")
	{
		logGroup.GET("/list", web.Login)

	}

	return router
}

// 根据env环境变量获取gin运行模式
func runMode() string {
	if config.Env() == constant.EnvPrd {
		return gin.ReleaseMode
	}
	return gin.DebugMode
}
