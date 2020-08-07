/*
@Time : 2020/7/15 11:04 下午
@Author : lucbine
@File : errcode.go
*/
package errcode

type Err struct {
	Code int
	Msg  string
}

//实现Error 接口
func (e *Err) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *Err {
	return &Err{
		Msg:  msg,
		Code: code,
	}
}
