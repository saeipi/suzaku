package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_auth"
)

func auth(group *gin.RouterGroup) {
	router := group.Group("auth")
	router.POST("register", api_auth.UserRegister)
}
