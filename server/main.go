/*
@Time : 2020/7/16 1:18 下午
@Author : lucbine
@File : main.go
*/
package main

import (
	"flag"
	"fmt"
	"hotwheels/server/app/router"
	"hotwheels/server/internal/config"
	"hotwheels/server/internal/core"
	"hotwheels/server/internal/logger"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/willas/overseer"
	"github.com/willas/overseer/fetcher"
)

func main() {
	serverAddr := fmt.Sprintf(":9555")
	file, err := exec.LookPath(os.Args[0])

	if err != nil {
		panic(err)
	}

	binPath, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	overseer.Run(overseer.Config{
		Program: prog,
		Address: serverAddr,
		Fetcher: &fetcher.File{Path: binPath},
		Debug:   true,
	})
}

func prog(state overseer.State) {
	env := flag.String("env", "prd", "set running env. options: local/qa/pre/prd")
	flag.Parse()
	fmt.Println("current use env : ", *env)

	//if err := service.Init(env); err != nil {
	//	panic(fmt.Sprintf("server init failed: %+v\n", err))
	//}

	//初始化trace
	//jeagertrace.InitJaeger(constant2.TraceProjectName, config.GetStringMap("trace.jaeger"))

	//开启pprof
	//pprof.Start()

	//初始化配置
	if err := config.InitConfig(*env); err != nil {
		panic(err)
	}

	//初始化日志
	if err := logger.InitLog(); err != nil {
		panic(err)
	}

	//初始化Db
	if err := core.InitGDB(); err != nil {
		panic(err)
	}

	apiHandler := router.InitRouter()

	s := &http.Server{
		Handler:        apiHandler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    config.GetDuration("config.server.readTimeout") * time.Millisecond,
		WriteTimeout:   config.GetDuration("config.server.writeTimeout") * time.Millisecond,
	}
	if err := s.Serve(state.Listener); err != nil {
		fmt.Println("server start failed", err)
	}

	//日志刷到磁盘
	if err := logger.Logger.Sync(); err != nil {
		fmt.Println("log sync err:", err.Error())
	}

	//等待以上结束
	time.Sleep(time.Second)
}
