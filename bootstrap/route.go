/*
 * @PackageName: bootstrap
 * @Description: 路由初始化
 * @Author: gabbymrh
 * @Date: 2021-10-03 14:32:52
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 14:32:52
 */

package bootstrap

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/http/middleware"
	rw "fast-go/routes/web"

	"github.com/gin-gonic/gin"
)

// SetupRoute 引导安装路由
func SetupRoute() *gin.Engine {
	router := gin.Default()

	// 资源文件路由
	router.Static("/resources", "./resources")
	// 找不到路由时的返回
	router.NoRoute(func(ctx *gin.Context) {
		http_response.QueryVoid(ctx, "请求路径不存在")
	})
	// Web 路由组注册
	rw.RegisterWebRoutes(router)

	// 全局中间件
	registerGlobalMiddleWare(router)

	return router
}

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Logger(),
		middleware.Cors(),
		middleware.Recovery(),
	)
}
