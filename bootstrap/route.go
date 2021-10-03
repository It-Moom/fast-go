/*
 * @PackageName: bootstrap
 * @Description: 路由启动文件
 * @Author: gabbymrh
 * @Date: 2021-10-03 14:32:52
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 14:32:52
 */

package bootstrap

import (
	"fast-go/app/http/middleware"
	rw "fast-go/routes/web"
	"github.com/gin-gonic/gin"
)

// SetupRoute 引导安装路由
func SetupRoute() *gin.Engine {
	router := gin.Default()

	// 资源文件路由
	router.Static("/resources", "./resources")
	// Web 路由组注册
	rw.RegisterWebRoutes(router)

	// 全局中间件
	// 跨域处理中间件
	router.Use(middleware.Cors())
	return router
}
