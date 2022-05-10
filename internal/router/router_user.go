package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_user"
)

func user(group *gin.RouterGroup) {
	router := group.Group("user")
	router.GET("self_info", api_user.SelfInfo)
	router.POST("edit_info", api_user.EditInfo)
}
