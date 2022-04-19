package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
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

func Error(ctx *gin.Context, err error, code int) {
	Err(ctx, err.Error(), int32(code))
}

func Err(ctx *gin.Context, err string, code int32) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}
