/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_task.go
*/
package model

import (
	"hotwheels/server/internal/core"
	"hotwheels/server/internal/errcode"
	"time"

	"github.com/jinzhu/gorm"
)

//记录结构体
type HtaskModel struct {
	Id              int       `gorm:"column:id" json:"id"`
	UserId          int       `gorm:"column:user_id" json:"user_id"`
	GroupId         int       `gorm:"column:group_id" json:"group_id"`
	TaskUser        string    `gorm:"column:task_user" json:"task_user"`
	TaskName        string    `gorm:"column:task_name" json:"task_name"`
	TaskType        string    `gorm:"column:task_type" json:"task_type"`
	Description     string    `gorm:"column:description" json:"description"`
	CronSpec        string    `gorm:"column:cron_spec" json:"cron_spec"`
	Concurrent      string    `gorm:"column:concurrent" json:"concurrent"`
	ConcurrentCount string    `gorm:"column:concurrent_count" json:"concurrent_count"`
	Command         string    `gorm:"column:command" json:"command"`
	Status          int       `gorm:"column:status" json:"status"`
	Notify          string    `gorm:"column:notify" json:"notify"`
	NotifyEmail     string    `gorm:"column:notify_email" json:"notify_email"`
	Timeout         int       `gorm:"column:timeout" json:"timeout"`
	ExecuteTimes    int       `gorm:"column:execute_times" json:"execute_times"`
	PrevTime        string    `gorm:"column:prev_time" json:"prev_time"`
	CreateTime      time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      time.Time `gorm:"column:update_time" json:"update_time"`
}

//构造函数
func NewHtaskModel() *HtaskModel {
	return &HtaskModel{}
}

//指定表
func (m *HtaskModel) TableName() string {
	return "h_task"
}

func (m *HtaskModel) Search(where string, bind []interface{}, offset int, limit int, orderBy string) (result []HtaskModel, err error) {
	err = core.DB.Table(m.TableName()).Where(where, bind...).Order(orderBy).Offset(offset).Limit(limit).Find(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, nil
	}
	return
}

//插入数据
func (m *HtaskModel) Add(data HtaskModel) *errcode.Err {
	data.CreateTime = time.Now()
	data.CreateTime = time.Now()
	err := core.DB.Table(m.TableName()).Create(data).Error
	return CreateErrorWrapper(err)
}

//更新数据
func (m *HtaskModel) update(id int, data HtaskModel) *errcode.Err {
	err := core.DB.Table(m.TableName()).Where("id = ?", id).Update(data).Error
	return UpdateErrorWrapper(err)
}
