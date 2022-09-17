/*
 * @PackageName: middleware
 * @Description: jwt授权中间件
 * @Author: gabbymrh
 * @Date: 2022-06-24 10:30:24
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-24 10:30:24
 */

package middleware

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/model/entity/user"
	"fast-go/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			http_response.TokenExpired(c, err.Error())
			return
		}

		// JWT 解析成功，设置用户信息
		// 传入的参数ID转为int
		_user := user.FindByID(cast.ToInt(claims.SubjectID))
		if _user.ID == 0 {
			http_response.TokenExpired(c, "用户不存在")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", _user.GetStringID())
		c.Set("current_user_name", _user.Name)
		c.Set("current_user", _user)

		c.Next()
	}
}
