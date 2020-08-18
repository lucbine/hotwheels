/*
@Time : 2020/7/11 10:04 上午
@Author : lucbine
@File : main.go
@version: v1.0
*/
package main

import (
	"flag"
	"fmt"
	"hotwheels/agent/internal/config"
	"hotwheels/agent/internal/core"
	"hotwheels/agent/internal/logger"
	"hotwheels/agent/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"
)

/*
   1、健康检查
   2、拉去配置信息
   3、日志上报
*/

func main() {
	//环境
	env := flag.String("env", "prd", "set running env. options: dev/qa/pre/prd")
	flag.Parse()
	fmt.Println("use env : ", *env)

	//初始化配置
	if err := config.InitConfig(*env); err != nil {
		panic(err)
	}

	//初始化日志
	if err := logger.InitLog(); err != nil {
		panic(err)
	}

	//初始化 redis
	if err := core.InitRedis(); err != nil {
		panic(err)
	}

	//初始化job
	c := cron.New(cron.WithSeconds())
	service.InitCronJob(c)

	//心跳检查 Heartbeat check
	hc := cron.New(cron.WithSeconds())
	hc.AddFunc("*/5 * * * * ?", func() {
		//service.NewNode().Check()
	})
	hc.Start()

	//监控通知信号
	exitChan := make(chan int)
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		fmt.Println("signal received")
		c.Stop()
		hc.Stop()
		exitChan <- 1
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan

}
