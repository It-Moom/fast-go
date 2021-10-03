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
	"fast-go/pkg/utils/generator"
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
func (j *JWT) CreateToken(userId string) (string, error) {

	// 过期时间处理
	expiresAt := time.Now().Add(time.Second * time.Duration(config.GetInt64("jwt.ttl"))).Unix()
	// 构造Payload
	claims := jwt.StandardClaims{
		ExpiresAt: expiresAt,                   // 过期时间
		Id:        generator.SimpleUUID(),      // Jti:token唯一标识符
		IssuedAt:  time.Now().Unix(),           // 颁发时间
		Issuer:    config.GetString("app.url"), // 颁发机构
		NotBefore: time.Now().Unix(),           // 生效时间
		Subject:   userId,                      // 主题
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	// token有误
	if err != nil {
		return nil, TokenInvalid
	}
	// token校验成功
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// RefreshToken 刷新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {

	// 解析token
	claims, err := NewJWT().ParseToken(tokenString)

	// token无效
	if err != nil {
		return response_message.TOKEN_EXPIRED, TokenInvalid
	}

	// 重新生成token
	return j.CreateToken(claims.Subject)

}
