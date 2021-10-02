/*
 * @PackageName: config
 * @Description: 应用配置
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:21:55
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:21:55
 */

package config

import "fast-go/pkg/config"

func init() {
	config.Add("app", config.StrMap{

		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "FastGo"),

		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),

		// 应用服务端口
		"port": config.Env("APP_PORT", "8080"),

		// APP密钥
		"key": config.Env("APP_KEY", "123456"),
	})
}
