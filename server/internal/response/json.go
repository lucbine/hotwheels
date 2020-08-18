/*
@Time : 2020/8/12 3:27 下午
@Author : lucbine
@File : json.go
*/
package response

import (
	"hotwheels/server/internal/errcode"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 发送接口 json 响应
func Json(ctx *gin.Context, err *errcode.Err, data interface{}, ext ...map[string]interface{}) {
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
	if len(ext) > 0 {
		for key, value := range ext[0] {
			resp[key] = value
		}
	}
	if err.Code < 0 {
		ctx.Set("errCode", err.Code)
		ctx.Set("errMsg", err.Msg)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
