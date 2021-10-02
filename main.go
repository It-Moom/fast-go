/*
 * @PackageName: main
 * @Description: 主入口
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:35:22
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:35:22
 */

package main

import (
	"fast-go/app/handler/http_response"
	"fast-go/config"
	c "fast-go/pkg/config"
	"fast-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func homeHandler(ctx *gin.Context) {
	http_response.RequestSuccess(ctx)
}

func main() {

	r := gin.Default()
	r.GET("/", homeHandler)
	err := r.Run(":" + c.GetString("app.port"))
	if err != nil {
		logger.LogError(err)
	}

}
