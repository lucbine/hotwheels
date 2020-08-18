/*
@Time : 2020/7/14 9:29 下午
@Author : lucbine
@File : agent.go
*/
package api

import (
	"hotwheels/server/app/entity"
	"hotwheels/server/app/service"
	"hotwheels/server/internal/errcode"
	"hotwheels/server/internal/logger"
	"hotwheels/server/internal/response"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//上报执行结果
func Report(ctx *gin.Context) {
	var reqs = new(entity.ReportReq)
	err := ctx.ShouldBind(reqs)
	if err != nil {
		logger.Logger.Error("report ctx bind is failed", zap.Error(err))
		response.Json(ctx, errcode.InputParamsError, nil)
		return
	}

	errcodeErr := service.NewLogService(ctx).Report(reqs)
	response.Json(ctx, errcodeErr, nil)
}

//拉去任务列表
func JobList(ctx *gin.Context) {
	jobList := make([]map[string]interface{}, 0)

	jobList = append(jobList, map[string]interface{}{
		"id":               1,
		"name":             "topic",
		"user":             "qtt",
		"spec":             "*/5 * * * * ?",
		"command":          "cd /data/demo && go run main.go",
		"status":           0,
		"timeout":          0,
		"concurrent":       false,
		"concurrent_count": 1,
	})

	jobList = append(jobList, map[string]interface{}{
		"id":               2,
		"name":             "circle",
		"user":             "qtt",
		"command":          "cd /data0/www/htdocs/gomod/topics/app/console && _publish_dir/topics-console/bin/topics-console cron --name=circleOneLatest  --env=local",
		"spec":             "0 22 17 * * ?",
		"status":           0,
		"timeout":          0,
		"concurrent":       false,
		"concurrent_count": 1,
	})

	jobList = append(jobList, map[string]interface{}{
		"id":               3,
		"name":             "content",
		"user":             "qtt",
		"spec":             "0 23 17 * * ?",
		"command":          "cd /data0/www/htdocs/gomod/topics/app/console && _publish_dir/topics-console/bin/topics-console cron --name=contentRelationCircle  --env=local",
		"status":           0,
		"timeout":          0,
		"concurrent":       false,
		"concurrent_count": 1,
	})

	data := map[string]interface{}{
		"list": jobList,
	}
	response.Json(ctx, errcode.Success, data)
}

//上报脚本机器的节点信息
func Hc(ctx *gin.Context) {
	var reqs = new(entity.HcReq)
	err := ctx.ShouldBind(reqs)
	if err != nil {
		logger.Logger.Error("hc ctx bind is failed", zap.Error(err))
		response.Json(ctx, errcode.InputParamsError, nil)
		return
	}
	errcodeErr := service.NewNodeService(ctx).Hc(reqs)
	response.Json(ctx, errcodeErr, nil)
}
