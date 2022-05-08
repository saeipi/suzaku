package router

import (
	"github.com/gin-gonic/gin"
	"suzaku/internal/interface/api/api_group"
)

func group(group *gin.RouterGroup) {
	router := group.Group("group")
	router.POST("member_list", api_group.MemberList)
}
