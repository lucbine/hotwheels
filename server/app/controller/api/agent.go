/*
@Time : 2020/7/14 9:29 下午
@Author : lucbine
@File : agent.go
*/
package api

import "github.com/gin-gonic/gin"

//上报执行结果
func Report(ctx *gin.Context) {
	/*
		1、记录执行日志
		2、发送告警
	*/

}

//拉去任务列表
func JobList(ctx *gin.Context) {

}

//上报脚本机器的节点信息
func NodeInfo(ctx *gin.Context) {

}

//监控检测
func Heal(ctx *gin.Context) {

}
