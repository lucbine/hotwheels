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

func NewServerClient() *ServerClient {
	var client = &ServerClient{
		host: config.GetString("config.server.url"),
	}

	return client
}

//发送脚本执行结果
func (c *ServerClient) Result(log *entity.ExecResult) (err error) {
	//path := "/hotwheels/pushLog"
	//var params = map[string]interface{}{
	//	//"content_ids": util.JoinIntSlice(contentIds, ","),
	//}
	//var result interface{}
	//err = c.callService("GET", path, params, &result)
	return
}

//任务列表
func (c *ServerClient) JobList() (list []entity.Job, err error) {
	//path := "/hotwheels/pushLog"
	//var params = map[string]interface{}{
	//	//"content_ids": util.JoinIntSlice(contentIds, ","),
	//}
	//var result interface{}
	//err = c.callService("GET", path, params, &result)
	return
}

//调用服务
func (c *ServerClient) callService(method string, path string, params map[string]interface{}, result interface{}) (err error) {
	//构建请求
	var url = fmt.Sprintf("%s%s", c.host, path)
	var httpReq = httpcall.Req{
		Url:     url,
		Params:  params,
		TimeOut: 300 * time.Millisecond,
	}
	if strings.ToUpper(method) == "GET" {
		err = httpcall.Get(httpReq, result)
	} else if strings.ToUpper(method) == "POST" {
		err = httpcall.Post(httpReq, result)
	} else {
		return fmt.Errorf("method = %s is error! ", method)
	}
	if err != nil {
		logger.Logger.Error("content service call fail",
			zap.Error(err), zap.Any("params", params))
	}
	return err
}
