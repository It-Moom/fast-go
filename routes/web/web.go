/*
 * @PackageName: routes
 * @Description: web 路由组
 * @Author: gabbymrh
 * @Date: 2021-10-03 14:42:17
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 14:42:17
 */

package web

import (
	cw "fast-go/app/http/controller/web"
	"github.com/gin-gonic/gin"
)

// RegisterWebRoutes Web 路由组
func RegisterWebRoutes(r *gin.Engine) {

	// 默认首页路由组
	webIndexRouter := r.Group("/")
	{
		ic := new(cw.IndexController)
		// 默认首页
		webIndexRouter.GET("/", ic.Index)
	}

}
