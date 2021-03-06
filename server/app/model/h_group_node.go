/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_group.go
*/
package model

import (
	"hotwheels/server/internal/core"
	"time"

	"github.com/jinzhu/gorm"
)

//记录结构体
type HGroupNodeModel struct {
	Id          int       `gorm:"column:id" json:"id"`
	TaskId      int64     `gorm:"column:task_id" json:"task_id"`
	output      string    `gorm:"column:output" json:"output"`
	error       string    `gorm:"column:error" json:"error"`
	Status      int8      `gorm:"column:status" json:"status"`
	ProcessTime int       `gorm:"column:process_time" json:"process_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

//构造函数
func NewHGroupNodeModel() *HGroupNodeModel {
	return &HGroupNodeModel{}
}

//指定表
func (m *HGroupNodeModel) TableName() string {
	return "content"
}

func (m *HGroupNodeModel) Search(where string, bind []interface{}, offset int, limit int, orderBy string) (result []HtaskModel, err error) {
	err = core.DB.Table(m.TableName()).Where(where, bind...).Order(orderBy).Offset(offset).Limit(limit).Find(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, nil
	}
	return
}
