/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_task.go
*/
package model

import (
	"hotwheels/server/internal/core"
	"hotwheels/server/internal/logger"
	"time"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
)

//记录结构体
type HtaskLogModel struct {
	Id          int       `gorm:"column:id" json:"id"`
	TaskId      int64     `gorm:"column:task_id" json:"task_id"`
	Output      string    `gorm:"column:output" json:"output"`
	Error       string    `gorm:"column:error" json:"error"`
	Status      int8      `gorm:"column:status" json:"status"`
	ProcessTime int       `gorm:"column:process_time" json:"process_time"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`
}

func NewHtaskLogModel() *HtaskLogModel {
	return &HtaskLogModel{}
}

//指定表
func (m *HtaskLogModel) TableName() string {
	return "h_task_log"
}

func (m *HtaskLogModel) Search(where string, bind []interface{}, offset int, limit int, orderBy string) (result []HtaskModel, err error) {
	err = core.DB.Table(m.TableName()).Where(where, bind...).Order(orderBy).Offset(offset).Limit(limit).Find(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, nil
	}
	return
}

//新增任务日志
func (m *HtaskLogModel) Add(data HtaskLogModel) (newId int, err error) {
	//默认项
	time := time.Now()
	data.UpdateTime = time
	data.CreateTime = time
	err = core.DB.Table(m.TableName()).
		Create(&data).Error
	if err != nil {
		logger.Logger.Error("db_add is fail", zap.Any("data", data), zap.Error(err))
		return 0, err
	}

	return data.Id, nil
}
