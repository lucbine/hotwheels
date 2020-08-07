/*
@Time : 2020/7/15 11:13 下午
@Author : lucbine
@File : response.go
*/
package response

import (
	"hotwheels/server/internal/errcode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 发送接口 json 响应
func Json(ctx *gin.Context, err *errcode.Err, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	resp := map[string]interface{}{
		"code":        err.Code,
		"message":     err.Msg,
		"showErr":     0,
		"currentTime": time.Now().Unix(),
		"data":        data,
	}
	if err.Code < 0 {
		ctx.Set("errCode", err.Code)
		ctx.Set("errMsg", err.Msg)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
