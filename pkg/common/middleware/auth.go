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
		if config.Config.JwtAuth.IsDev {
			ctx.Set(constant.KeyUserId, constant.KeyUserAdminUserId)
			ctx.Set(constant.KeyUserPlatform, constant.KeyUserAdminPlatform)
			return
		}
		token, err = jwt_auth.ParseJwtFromCookie(ctx)
		if err != nil {
			ctx.Abort()
			http.Error(ctx, err, http.ErrorCodeHttpJwtTokenErr)
			return
		}
		claims := jwt.MapClaims{}
		for key, value := range token.Claims.(jwt.MapClaims) {
			claims[key] = value
		}
		if _, ok = claims[constant.KeyUserId]; ok == false || utils.ToString(claims[constant.KeyUserId]) == "" {
			ctx.Abort()
			http.Error(ctx, http.ErrorHttpUseridDoesNotExist, http.ErrorCodeHttpUseridDoesNotExist)
			return
		}
		if _, ok = claims[constant.KeyUserPlatform]; ok == false || utils.TryToInt(claims[constant.KeyUserPlatform]) == 0 {
			ctx.Abort()
			http.Error(ctx, http.ErrorHttpUseridDoesNotExist, http.ErrorCodeHttpUseridDoesNotExist)
			return
		}
		ctx.Set(constant.KeyUserId, claims[constant.KeyUserId])
		ctx.Set(constant.KeyUserPlatform, claims[constant.KeyUserPlatform])
	}
}
