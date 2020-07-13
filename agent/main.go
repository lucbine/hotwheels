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
	"hotwheels/agent/internal/logger"
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

}
