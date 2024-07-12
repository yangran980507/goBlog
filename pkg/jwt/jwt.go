// Package jwt JWT 鉴权
package jwt

import (
	"blog/global"
	"blog/pkg/app"
	"blog/pkg/errcode"
	"blog/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// JWT 对象
type JWT struct {
	SecretKey  []byte // 密钥
	MaxRefresh time.Duration
}

// UserInfo 用户信息 私有声明
type UserInfo struct {
	UserID        uint   `json:"user_id"`
	UserLoginName string `json:"user_login_name"`
}

// CustomJWTClaims 自定义 payload 信息
type CustomJWTClaims struct {
	UserInfo
	ExpireAt int64
	jwtpkg.RegisteredClaims
}

// NewJWT 创建 JWT 实例
func NewJWT() *JWT {

	return &JWT{
		SecretKey:  []byte(global.AppSetting.JWTSecretKey),
		MaxRefresh: time.Duration(global.AppSetting.JWTExpireTime) * time.Minute,
	}
}

// IssueToken 签发 token
func (jwt *JWT) IssueToken(userinfo UserInfo) string {

	expire := jwt.expireTime()
	claims := &CustomJWTClaims{
		UserInfo: UserInfo{
			UserID:        userinfo.UserID,        // 用户ID
			UserLoginName: userinfo.UserLoginName, // 用户登陆名
		},
		ExpireAt: expire.Unix(),
		RegisteredClaims: jwtpkg.RegisteredClaims{
			Issuer:    global.AppSetting.Name,            // 签发者
			IssuedAt:  jwtpkg.NewNumericDate(time.Now()), // 签发日期
			NotBefore: jwtpkg.NewNumericDate(time.Now()), // 签发生效日期
			ExpiresAt: jwtpkg.NewNumericDate(expire),     // 签发过期日期
		},
	}

	token, err := jwt.generateToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// 生成 Token
func (jwt *JWT) generateToken(claims *CustomJWTClaims) (string, error) {
	return jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims).SignedString(jwt.SecretKey)
}

// 获取过期时间
func (jwt *JWT) expireTime() time.Time {
	timeNow := time.Now()
	var expireTime int64
	if app.IsLocal() {
		// 本地环境过期时间两个月
		expireTime = global.AppSetting.JWTMaxExpireTime
	} else {
		// 线上环境过期时间2小时
		expireTime = global.AppSetting.JWTExpireTime
	}
	expire := time.Duration(expireTime) * time.Minute
	return timeNow.Add(expire)
}

// ParseToken 解析 Token
func (jwt *JWT) ParseToken(c *gin.Context) (*CustomJWTClaims, error) {

	// 从 Request.Header 中读取 Token
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return &CustomJWTClaims{}, err
	}

	// 解析 TokenString
	token, err := jwt.parseTokenString(tokenString)

	// 解析出错
	if err != nil {
		return nil, err
	}

	// 解析出的数据与 CustomJWTClaims 结构校验
	if token != nil {
		if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// 获取 Request.Header 中的 Token
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	// 请求头中获取 token
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		// 字段为空
		return "", errors.New("需要登陆才能访问")
	}

	tokenSlice := strings.SplitN(tokenString, " ", 2)
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		// 格式错误
		return "", errors.New("token错误")
	}
	return tokenSlice[1], nil
}

// jwt.ParseWithClaims 解析 tokenString
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &CustomJWTClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SecretKey, nil
	})
}

// RefreshToken 刷新 token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 解析 token
	claims, err := jwt.ParseToken(c)
	if err != nil {
		return "", err
	}

	t := time.Now().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt.Unix() > t {
		claims.RegisteredClaims.ExpiresAt = jwtpkg.NewNumericDate(jwt.expireTime())
		return jwt.generateToken(claims)
	}
	return "", errcode.ErrTokenTimeOut
}
