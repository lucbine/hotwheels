/*
@Time : 2020/7/15 11:04 下午
@Author : lucbine
@File : errcode.go
*/
package errcode

//通用
var (
	Success          = NewError(0, "成功")
	ErrorFail        = NewError(-1, "操作失败")
	InputParamsError = NewError(-2, "参数输入错误")
	LoginError       = NewError(-4, "请登录后操作")
)

//db 错误码
const (
	DbInsertErrorCode = -100
	DbUpdateErrorCode = -101
	DbSelectErrorCode = -102
)

//参数校验
var (
	CronFormatError = NewError(-1000, "cron表达式无效")
)
