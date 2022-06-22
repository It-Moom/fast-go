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
	btsConfig "fast-go/config"
	"fast-go/pkg/config"
	pc "fast-go/pkg/config"
	"flag"
)

func init() {
	// 初始化配置信息
	btsConfig.Initialize()
}

// 入口函数
func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .yml 文件，如 --env=testing 加载的是 .yml.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 初始化 SQL
	bootstrap.SetupDB()
	// 初始化服务及路由
	bootstrap.SetupRoute().Run(":" + pc.GetString("app.port"))
}
