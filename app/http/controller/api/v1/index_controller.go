/*
 * @PackageName: v1
 * @Description: v1版本示例api控制器
 * @Author: gabbymrh
 * @Date: 2022-06-23 21:10:29
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 21:10:29
 */

package v1

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/http/controller"
	"github.com/gin-gonic/gin"
)

// IndexController 示例API控制器结构体
type IndexController struct {
	controller.BaseController
}

// Index 首页默认方法
func (ic *IndexController) Index(ctx *gin.Context) {
	http_response.RequestSuccess(ctx, "欢迎使用fast-go-api v1")
}
