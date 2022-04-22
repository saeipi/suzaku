package router_api

import (
	"github.com/gin-gonic/gin"
	"suzaku/pkg/common/middleware"
)

func RegisterRouter(engine *gin.Engine) {
	publicGroup := engine.Group("open")
	registerPublicRoutes(publicGroup)

	privateGroup := engine.Group("api")
	registerPrivateRouter(privateGroup)
}

func registerPublicRoutes(group *gin.RouterGroup) {
	auth(group)
}

func registerPrivateRouter(group *gin.RouterGroup) {
	// 授权验证
	group.Use(middleware.JwtAuth())
	user(group)
	friend(group)
	chat(group)
	minio(group)
}
