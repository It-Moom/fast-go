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
	"fast-go/app/http/middleware"
	"fast-go/pkg/config"
	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes Api 路由组
func RegisterApiRoutes(r *gin.Engine) {

	// API路由组
	var apiV1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		apiV1 = r.Group("/api/v1")
	} else {
		apiV1 = r.Group("/v1")
	}

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	apiV1.Use(middleware.LimitIP("200-H"))
	{
		// API欢迎路由
		ic := new(cav1.IndexController)
		// 默认首页
		apiV1.GET("/", ic.Index)

		// 用户路由组
		userGroup := apiV1.Group("/user")
		{
			uc := new(cav1.UserController)
			// 用户数据
			userGroup.GET("/info/:id", uc.ShowOne)
			// 用户列表
			userGroup.GET("/list", uc.GetAll)
			// 新增用户
			userGroup.POST("/create", uc.Create)
			// 更新用户
			userGroup.PUT("/update/:id", uc.Update)
			// 删除用户
			userGroup.DELETE("/delete/:id", uc.Delete)
		}
	}

}
