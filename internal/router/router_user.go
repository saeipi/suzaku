package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_user"
)

func user(group *gin.RouterGroup) {
	router := group.Group("user")
	router.GET("user_info", api_user.UserInfo)
	router.POST("edit_info", api_user.EditInfo)
}
