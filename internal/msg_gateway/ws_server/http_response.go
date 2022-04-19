package ws_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpSuccess(ctx *gin.Context, data interface{}) {
	if data == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func httpError(ctx *gin.Context, err error, code int) {
	httpErr(ctx, err.Error(), int32(code))
}

func httpErr(ctx *gin.Context, err string, code int32) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}
