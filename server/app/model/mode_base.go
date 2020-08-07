/*
@Time : 2020/7/14 9:56 下午
@Author : lucbine
@File : h_task.go
*/
package model

import "hotwheels/server/internal/errcode"

func CreateErrorWrapper(err error) *errcode.Err {
	if err != nil {
		return errcode.NewError(errcode.DbInsertErrorCode, err.Error())
	}
	return errcode.Success
}

func UpdateErrorWrapper(err error) *errcode.Err {
	if err != nil {
		return errcode.NewError(errcode.DbUpdateErrorCode, err.Error())
	}
	return errcode.Success
}
