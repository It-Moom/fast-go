/*
 * @PackageName: web
 * @Description: 首页控制器
 * @Author: gabbymrh
 * @Date: 2021-10-03 14:34:32
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 14:34:32
 */

package web

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/http/controller"
	"github.com/gin-gonic/gin"
)

// IndexController 默认首页控制器结构体
type IndexController struct {
	controller.BaseController
}

// Index 首页默认方法
func (ic *IndexController) Index(ctx *gin.Context) {
	http_response.RequestSuccess(ctx, "欢迎使用fast-go")
}
