package utils

import (
	"github.com/gin-gonic/gin"
	"suzaku/pkg/constant"
	"suzaku/pkg/http"
)

func RequestIdentity(c *gin.Context) (uidStr string, platformVal int32, ok bool) {
	var (
		userId   interface{}
		platform interface{}
		isExist  bool
	)
	userId, isExist = c.Get(constant.KeyUserID)
	if isExist == false {
		http.Error(c, http.ErrorHttpUserIdDoesNotExist, http.ErrorCodeHttpUserIdDoesNotExist)
		return
	}
	uidStr = ToString(userId)
	if uidStr == "" {
		http.Error(c, http.ErrorHttpUserIdDoesNotExist, http.ErrorCodeHttpUserIdDoesNotExist)
		return
	}
	platform, isExist = c.Get(constant.KeyUserPlatformID)
	if isExist == false {
		http.Error(c, http.ErrorHttpPlatformIdDoesNotExist, http.ErrorCodeHttpPlatformIdDoesNotExist)
		return
	}
	platformVal = int32(TryToInt(platform))
	if platformVal == 0 {
		http.Error(c, http.ErrorHttpPlatformIdDoesNotExist, http.ErrorCodeHttpPlatformIdDoesNotExist)
		return
	}
	ok = true
	return
}
