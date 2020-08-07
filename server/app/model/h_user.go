/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_task.go
*/
package model

import (
	"hotwheels/server/internal/core"
	"time"

	"github.com/jinzhu/gorm"
)

//记录结构体
type HuserModel struct {
	Id         int       `gorm:"column:id" json:"id"`
	UserName   int64     `gorm:"column:user_name" json:"user_name"`
	Email      string    `gorm:"column:email" json:"email"`
	Password   string    `gorm:"column:password" json:"password"`
	LastLogin  int       `gorm:"column:last_login" json:"last_login"`
	LastIp     int       `gorm:"column:last_ip" json:"last_ip"`
	Status     int8      `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
}

//构造函数
func NewHuserModel() *HuserModel {
	return &HuserModel{}
}

//指定表
func (m *HuserModel) TableName() string {
	return "content"
}

func (m *HuserModel) Search(where string, bind []interface{}, offset int, limit int, orderBy string) (result []HtaskModel, err error) {
	err = core.DB.Table(m.TableName()).Where(where, bind...).Order(orderBy).Offset(offset).Limit(limit).Find(&result).Error
	if gorm.IsRecordNotFoundError(err) {
		return result, nil
	}
	return
}
