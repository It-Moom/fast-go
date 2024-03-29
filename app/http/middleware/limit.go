/*
 * @PackageName: middleware
 * @Description: 限流中间件
 * @Author: gabbymrh
 * @Date: 2022-06-22 20:37:36
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 20:37:36
 */

package middleware

import (
	"fast-go/app/handler/http_response"
	"fast-go/pkg/app"
	"fast-go/pkg/limiter"
	"fast-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// LimitIP 全局限流中间件，针对 IP 进行限流
// limit 为格式化字符串，如 "5-S" ，示例:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
//
func LimitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}

	return func(ctx *gin.Context) {
		// 针对 IP 限流
		key := limiter.GetKeyIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

// LimitPerRoute 限流中间件，用在单独的路由中
func LimitPerRoute(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "1000000-H"
	}
	return func(ctx *gin.Context) {

		// 针对单个路由，增加访问次数
		ctx.Set("limiter-once", false)

		// 针对 IP + 路由进行限流
		key := limiter.GetKeyRouteWithIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

// limitHandler 限流处理
func limitHandler(ctx *gin.Context, key string, limit string) bool {

	// 获取超额的情况
	rate, err := limiter.CheckRate(ctx, key, limit)
	if err != nil {
		logger.LogIf(err)
		http_response.RequestError(ctx, "服务器内部错误")
		return false
	}

	// ---- 设置标头信息-----
	// X-RateLimit-Limit :10000 最大访问次数
	// X-RateLimit-Remaining :9993 剩余的访问次数
	// X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
	ctx.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	ctx.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	ctx.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	// 超额
	if rate.Reached {
		http_response.TooManyRequest(ctx, "接口请求太频繁")
		return false
	}

	return true
}
