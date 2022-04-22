package jwt_auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	JwtIssuer               = "tKh3fMwWepCds9jbTuJ9IGYJYbPIZnj"
	JwtTokenKey             = "suzaku_jwt_token_2021"
	JwtTokenExpireIn        = 3600 * 24 * 6
	JwtRefreshTokenExpireIn = 3600 * 24 * 30 // 刷新token的时长
	JwtRefreshTokenKey      = "suzaku_jwt_refresh_token_2021"
)

func CreateJwtToken(userId string, platform int32) (tokenString string, expireIn int64) {
	var (
		token  *jwt.Token
		claims jwt.MapClaims
		err    error
	)
	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)

	expireIn = JwtTokenExpireIn
	claims["iss"] = JwtIssuer
	claims["exp"] = time.Now().Add(time.Duration(expireIn) * time.Second).Unix()
	claims["orig_iat"] = time.Now().Unix()
	claims["user_id"] = userId
	claims["platform_id"] = platform
	tokenString, err = token.SignedString([]byte(JwtTokenKey))
	if err != nil {
		expireIn = -1
		return
	}
	return
}

//CreateJwtRefreshToken 创建刷新token
func CreateJwtRefreshToken(userId string, platform int32) (tokenString string, expireIn int64) {
	var (
		token  *jwt.Token
		claims jwt.MapClaims
		err    error
	)
	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)
	expireIn = JwtRefreshTokenExpireIn
	claims["iss"] = JwtIssuer
	claims["exp"] = time.Now().Add(time.Duration(expireIn) * time.Second).Unix()
	claims["orig_iat"] = time.Now().Unix()
	claims["user_id"] = userId
	claims["platform_id"] = platform
	tokenString, err = token.SignedString([]byte(JwtRefreshTokenKey))
	if err != nil {
		expireIn = -1
	}
	return
}

//ParseJwtRefreshToken解析刷新Token
func ParseJwtRefreshToken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtRefreshTokenKey), nil
	})
	return
}

func ParseJwtFromHeader(ctx *gin.Context) (res *jwt.Token, err error) {
	res, err = request.ParseFromRequest(ctx.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtTokenKey), nil
		})
	if err == request.ErrNoTokenInRequest {
		token := ctx.Query("token")
		res, err = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(JwtTokenKey), nil
		})
	}
	return
}

func ParseJwtFromCookie(ctx *gin.Context) (*jwt.Token, error) {
	token := ctx.Query("token")
	cookie, _ := ctx.Cookie("jwt")
	tokenStr := cookie
	if token != "" {
		tokenStr = token
	}
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtTokenKey), nil
	})
}

func GetFeishuAccessToken(userId, appId, appSecret string) (accessToken string) {
	//accessToken, _ = redis.RedisGet(fmt.Sprintf(redis.RedisKeyFeishuAccessToken, userId))
	//if accessToken != "" {
	//	return
	//}
	//
	//refreshToken, _ := redis.RedisGet(fmt.Sprintf(redis.RedisKeyFeishuRefreshToken, userId))
	//if refreshToken == "" {
	//	return
	//}
	//
	//app, err := request_feishu.NewApp(appId, appSecret)
	//if err != nil {
	//	return
	//}
	//
	//user := &request_feishu.User{
	//	RefreshToken: refreshToken,
	//}
	//
	//resp, err := user.RefreshAccessToken(app.AppAccessToken)
	//if err != nil {
	//	return
	//}
	//redis.RedisSet(
	//	fmt.Sprintf(redis.RedisKeyFeishuAccessToken, userId),
	//	resp.Data.AccessToken,
	//	int(resp.Data.ExpiresIn*1000))
	//
	//redis.RedisSet(
	//	fmt.Sprintf(redis.RedisKeyFeishuRefreshToken, userId),
	//	resp.Data.RefreshToken,
	//	int(resp.Data.RefreshExpiresIn*1000))
	//
	//accessToken = resp.Data.AccessToken
	return
}

func ParseJwtFromToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtTokenKey), nil
	})
}
