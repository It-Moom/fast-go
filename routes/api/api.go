/*
 * @PackageName: routes
 * @Description: 路由组
 * @Author: gabbymrh
 * @Date: 2021-10-03 14:29:42
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 14:29:42
 */

package api

import (
	cav1 "fast-go/app/http/controller/api/v1"
	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes Api 路由组
func RegisterApiRoutes(r *gin.Engine) {

	// API路由组
	webIndexRouter := r.Group("/api/v1")
	{
		ic := new(cav1.IndexController)
		// 默认首页
		webIndexRouter.GET("/", ic.Index)
	}

}
