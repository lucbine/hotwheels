/*
@Time : 2020/7/13 9:49 下午
@Author : lucbine
@File : email.go
*/
package msg

import "hotwheels/agent/entity"

type Msg interface {
	Send(entity.Notice) error //发送消息
}
