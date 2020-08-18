/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_group.go
*/
package model

import (
	"hotwheels/server/internal/core"
	"hotwheels/server/internal/errcode"
	"time"

	"github.com/jinzhu/gorm"
)

//记录结构体
type HNodeModel struct {
	Id          int       `gorm:"column:id" json:"id"`
	IP          string    `gorm:"column:ip" json:"ip"`
	HostName    string    `gorm:"column:host_name" json:"host_name"`
	CpuCount    int       `gorm:"column:cpu_count" json:"cpu_count"`
	Os          string    `gorm:"column:os" json:"os"`
	CpuUsage    float64   `gorm:"column:cpu_usage" json:"cpu_usage"`
	MemorySize  uint64    `gorm:"column:memory_size" json:"memory_size"`
	MemoryUsage float64   `gorm:"column:memory_usage" json:"memory_usage"`
	Status      int8      `gorm:"column:status" json:"status"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

//构造函数
func NewHNodeModel() *HNodeModel {
	return &HNodeModel{}
}

//指定表
func (m *HNodeModel) TableName() string {
	return "h_node"
}

func (m *HNodeModel) Search(where string, bind []interface{}, offset int, limit int, orderBy string) (result []HtaskModel, err error) {
	err = core.DB.Table(m.TableName()).Where(where, bind...).Order(orderBy).Offset(offset).Limit(limit).Find(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, nil
	}
	return
}

//插入数据
func (m *HNodeModel) Add(data HNodeModel) *errcode.Err {
	time := time.Now()
	data.CreateTime = time
	data.UpdateTime = time
	err := core.DB.Table(m.TableName()).Create(&data).Error
	return CreateErrorWrapper(err)
}

//更新数据

func (m *HNodeModel) Update(id int, data HNodeModel) *errcode.Err {
	time := time.Now()
	data.CreateTime = time
	data.UpdateTime = time
	err := core.DB.Table(m.TableName()).Where("id = ?", id).Update(data).Error
	return UpdateErrorWrapper(err)
}

//通过ip 查询
func (m *HNodeModel) GetByIp(ip string) (result HNodeModel, er *errcode.Err) {
	err := core.DB.Table(m.TableName()).Where("ip = ?", ip).First(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, SelectErrorWrapper(nil)
	}
	return result, SelectErrorWrapper(err)
}
