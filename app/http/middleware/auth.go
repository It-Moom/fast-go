/**
 * @PackageName: middleware
 * @Description: 鉴权中间件
 * @Author: Casso
 * @Date: 2021-10-03 15:52
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-03 17:04:58
 */

package middleware

import (
	"fast-go/app/handler/http_response"
	auth "fast-go/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
)

// JwtAuth 中间件，检查token
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			http_response.TokenExpired(ctx)
			ctx.Abort()
			return
		}

		// // 自定义token拼接格式 -- 前端请求头需添加同样格式(请求拦截器里面设置)
		// var token = ""
		// if len(strings.Split(tokenStr, " ")) > 1 && strings.Split(tokenStr, " ")[0] == "GOJWT" {
		// 	token = strings.Split(tokenStr, " ")[1]
		// }
		// if token == "" {
		// 	http_response.TokenExpired(ctx)
		// 	ctx.Abort()
		// 	return
		// }

		j := auth.NewJWT()
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			http_response.TokenExpired(ctx)
			ctx.Abort()
		}

		// 继续交由下一个路由处理,并将解析出的信息传递下去
		ctx.Set("userId", claims.Subject)
		ctx.Next()
	}
}
