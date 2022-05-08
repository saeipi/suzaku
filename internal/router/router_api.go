package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/pkg/common/middleware"
)

func Register(engine *gin.Engine) {
	publicRouter := engine.Group("open")
	registerPublicRoutes(publicRouter)

	privateRouter := engine.Group("api")
	registerPrivateRouter(privateRouter)
}

func registerPublicRoutes(router *gin.RouterGroup) {
	auth(router)
}

func registerPrivateRouter(router *gin.RouterGroup) {
	// 授权验证
	router.Use(middleware.JwtAuth())
	user(router)
	friend(router)
	chat(router)
	minio(router)
	group(router)
}
