/*
@Time : 2020/7/15 10:44 下午
@Author : lucbine
@File : node.go
*/
package service

import (
	"hotwheels/server/app/constant"
	"hotwheels/server/app/entity"
	"hotwheels/server/app/model"
	"hotwheels/server/internal/errcode"

	"github.com/gin-gonic/gin"
)

type NodeService struct {
	ctx       *gin.Context
	nodeModel *model.HNodeModel
}

func NewNodeService(c *gin.Context) *NodeService {
	return &NodeService{
		ctx:       c,
		nodeModel: model.NewHNodeModel(),
	}
}

//节点监控检测
func (ns *NodeService) Hc(hcReq *entity.HcReq) *errcode.Err {
	//先检查是否有用户状态
	var (
		node model.HNodeModel
		err  *errcode.Err
	)

	node, err = ns.nodeModel.GetByIp(hcReq.IP)
	data := model.HNodeModel{
		IP:          hcReq.IP,
		HostName:    hcReq.HostName,
		CpuCount:    hcReq.CpuCount,
		Os:          hcReq.Os,
		CpuUsage:    hcReq.CpuUsage,
		MemorySize:  hcReq.MemorySize,
		MemoryUsage: hcReq.MemoryUsage,
		Status:      constant.NodeStatusOnline,
	}

	if node.Id != 0 {
		//更新信息
		err = ns.nodeModel.Update(node.Id, data)

	} else { //插入信息
		err = ns.nodeModel.Add(data)
	}
	return err
}

//添加定时任务
func (ns *NodeService) add() *errcode.Err {

	return errcode.Success
}

//定时任务列表
func (ns *NodeService) List() *errcode.Err {

	return errcode.Success
}

//编辑任务
func (ns *NodeService) Edit() *errcode.Err {

	return errcode.Success
}

//任务统计
func (ns *NodeService) Stat() *errcode.Err {

	return errcode.Success
}
