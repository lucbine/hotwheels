/*
@Time : 2020/7/13 3:17 下午
@Author : lucbine
@File : jobs.go
*/
package service

import (
	"hotwheels/agent/internal/logger"
	"hotwheels/agent/internal/util"
	"os"
	"runtime"
	"time"

	"go.uber.org/zap"

	"github.com/shirou/gopsutil/cpu"

	"github.com/shirou/gopsutil/mem"
)

//节点信息
type Node struct { //任务
	IP          string  //内网IP
	HostName    string  //主机名称
	CpuCount    int     //cpu个数
	Os          string  //系统名称
	CpuUsage    float64 //cpu 使用率
	MemorySize  uint64  //内存大小
	MemoryUsage float64 //内存使用率
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Check() error {
	n.HostName, _ = os.Hostname()
	n.CpuCount = runtime.NumCPU()
	n.Os = runtime.GOOS
	ip, _ := util.GetExternalIP()
	n.IP = ip.String()
	n.CpuUsage = GetCpuPercent()
	n.MemoryUsage = GetMemPercent()
	n.MemorySize = GetMem()
	err := NewServerClient().Hc(n)
	logger.Logger.Info("hc", zap.String("ip", n.IP), zap.String("host_name", n.HostName), zap.Int("cpu_count", n.CpuCount),
		zap.String("os", n.Os), zap.Float64("cpu_usage", n.CpuUsage), zap.Uint64("memory_size", n.MemorySize), zap.Float64("memory_usage", n.MemoryUsage))
	return err
}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetMem() uint64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.Total
}
