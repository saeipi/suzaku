package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"suzaku/pkg/common/config"
	"suzaku/pkg/common/jwt_auth"
	"suzaku/pkg/constant"
	"suzaku/pkg/http"
	"suzaku/pkg/utils"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			err   error
			token *jwt.Token
			ok    bool
		)
		switch config.Config.JwtAuth.AuthMethod {
		case constant.HttpKeyAuthMethodCookie:
			token, err = jwt_auth.ParseJwtFromCookie(ctx)
		case constant.HttpKeyAuthMethodHeader:
			token, err = jwt_auth.ParseJwtFromHeader(ctx)
		default:
			return
		}
		/*
			if config.Config.JwtAuth.IsDev {
				ctx.Set(constant.KeyUserID, constant.KeyUserAdminUserId)
				ctx.Set(constant.KeyUserPlatformID, constant.KeyUserAdminPlatform)
			}*/
		if err != nil {
			ctx.Abort()
			http.Error(ctx, err, http.ErrorCodeHttpJwtTokenErr)
			return
		}
		claims := jwt.MapClaims{}
		for key, value := range token.Claims.(jwt.MapClaims) {
			claims[key] = value
		}
		if _, ok = claims[constant.KeyUserID]; ok == false {
			ctx.Abort()
			http.Error(ctx, http.ErrorHttpUserIdDoesNotExist, http.ErrorCodeHttpUserIdDoesNotExist)
			return
		}
		if _, ok = claims[constant.KeyUserPlatformID]; ok == false || utils.TryToInt(claims[constant.KeyUserPlatformID]) == 0 {
			ctx.Abort()
			http.Error(ctx, http.ErrorHttpPlatformIdDoesNotExist, http.ErrorCodeHttpPlatformIdDoesNotExist)
			return
		}
		ctx.Set(constant.KeyUserID, claims[constant.KeyUserID])
		ctx.Set(constant.KeyUserPlatformID, claims[constant.KeyUserPlatformID])
	}
}
