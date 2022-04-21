package router_api

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_minio"
)

func minio(group *gin.RouterGroup) {
	router := group.Group("minio")
	router.POST("upload", api_minio.UploadFile)
}
