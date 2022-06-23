/*
 * @PackageName: middleware
 * @Description: 强制 UserAgent 中间件
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:52:49
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:52:49
 */

package middleware

import (
	"fast-go/app/handler/http_response"
	"github.com/gin-gonic/gin"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 User-Agent 标头信息:User-Agent 携带客户端的信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			http_response.RequestError(c, "请求头中缺少 User-Agent 标头,请检查请求头")
			return
		}

		c.Next()
	}
}
