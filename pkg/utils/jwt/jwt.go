/**
 * @PackageName: jwt
 * @Description: jwt相关
 * @Author: Casso-Wong
 * @Date: 2021-10-03 16:42:59
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 18:49:21
 */

package jwt

import (
	"errors"
	"fast-go/app/constant/response_message"
	"fast-go/pkg/config"
	"github.com/golang-jwt/jwt"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

var (
	TokenInvalid error  = errors.New(response_message.TOKEN_EXPIRED)
	SignKey      string = config.GetString("jwt.secret")
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Id int `json:"id"` // subject id
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// 设置自定义token过期时间
	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(config.GetInt64("jwt.ttl"))).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// token有误
	if err != nil {
		return nil, TokenInvalid
	}
	// token校验成功
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	// 校验token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// token无效
	if err != nil {
		return response_message.TOKEN_EXPIRED, TokenInvalid
	}

	// token校验有效则返回新token
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(time.Duration(config.GetInt64("jwt.ttl"))).Unix()
		return j.CreateToken(*claims)
	}
	return response_message.TOKEN_EXPIRED, TokenInvalid
}
