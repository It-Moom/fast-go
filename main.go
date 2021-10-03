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
	"fast-go/bootstrap"
	"fast-go/config"
	pc "fast-go/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

// 入口函数
func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化服务及路由
	bootstrap.SetupRoute().Run(":" + pc.GetString("app.port"))
}
