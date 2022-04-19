package router_api

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_friend"
)

func friend(group *gin.RouterGroup) {
	router := group.Group("friend")
	router.POST("add_friend", api_friend.AddFriend)
}
