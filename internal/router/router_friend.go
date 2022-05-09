package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_friend"
)

func friend(group *gin.RouterGroup) {
	router := group.Group("friend")
	router.POST("add_friend", api_friend.AddFriend)
	router.GET("friend_request_list", api_friend.FriendRequestList)
	router.POST("handle_friend_request", api_friend.HandleFriendRequest)
}
