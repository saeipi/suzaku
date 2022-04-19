package router_api

import "github.com/gin-gonic/gin"

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
	user(group)
	friend(group)
	chat(group)
}
