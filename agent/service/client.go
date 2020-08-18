/*
@Time : 2020/7/13 9:05 下午
@Author : lucbine
@File : client.go
*/
package service

import (
	"fmt"
	"hotwheels/agent/entity"
	"hotwheels/agent/internal/config"
	"hotwheels/agent/internal/httpcall"
	"hotwheels/agent/internal/logger"
	"strings"
	"time"

	"go.uber.org/zap"
)

/*
	服务端api 调用
*/
type ServerClient struct {
	host string
}

type JobListResp struct {
	entity.BaseResponse
	Data struct {
		List []entity.Task `json:"list"`
	} `json:"data"`
}

func NewServerClient() *ServerClient {
	var client = &ServerClient{
		host: config.GetString("config.server.url"),
	}

	return client
}

//监控检查
func (c *ServerClient) Hc(node *Node) (err error) {
	path := "/api/agent/hc"
	var params = map[string]interface{}{
		"ip":           node.IP,
		"host_name":    node.HostName,
		"cpu_count":    node.CpuCount,
		"os":           node.Os,
		"cpu_usage":    node.CpuUsage,
		"memory_size":  node.MemorySize,
		"memory_usage": node.MemoryUsage,
	}
	var result interface{}
	err = c.callService("POST", path, params, &result)
	return
}

//发送脚本执行结果
func (c *ServerClient) Result(log *entity.ExecResult) (err error) {
	path := "/api/agent/report"
	var params = map[string]interface{}{
		"task_id":      log.TaskId,
		"output":       log.Output,
		"error":        log.Error,
		"process_time": log.ProcessTime,
		"create_time":  log.CreateTime,
		"status":       log.Status,
	}
	var result interface{}
	err = c.callService("POST", path, params, &result)
	return
}

//任务列表
func (c *ServerClient) TaskList() (list []entity.Task, err error) {
	path := "/api/agent/jobList"
	var params = map[string]interface{}{}
	var res JobListResp
	err = c.callService("GET", path, params, &res)
	return res.Data.List, err
}

//调用服务
func (c *ServerClient) callService(method string, path string, params map[string]interface{}, result interface{}) (err error) {
	//构建请求
	var url = fmt.Sprintf("%s%s", c.host, path)
	var httpReq = httpcall.Req{
		Url:     url,
		Params:  params,
		TimeOut: 300 * time.Millisecond,
		Header:  make(map[string]string),
	}
	if strings.ToUpper(method) == "GET" {
		err = httpcall.Get(httpReq, result)
	} else if strings.ToUpper(method) == "POST" {
		httpReq.Header["Content-Type"] = "application/x-www-form-urlencoded"
		err = httpcall.PostForm(httpReq, result)
	} else {
		return fmt.Errorf("method = %s is error! ", method)
	}
	//logger.Logger.Info("callService", zap.String("method", method), zap.String("url", url),
	//	zap.Any("params", params), zap.Any("result", result), zap.Error(err))
	if err != nil {
		logger.Logger.Error("content service call fail",
			zap.Error(err), zap.Any("params", params))
	}
	return err
}
