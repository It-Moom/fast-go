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
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称，暂时没有使用到
			"name": config.Env("app.name", "FastGo"),

			// 当前环境，用以区分多环境
			"env": config.Env("app.env", "production"),

			// 是否进入调试模式
			"debug": config.Env("app.debug", false),

			// App url
			"url": config.Env("app.url", "http://localhost"),

			// 应用服务端口
			"port": config.Env("app.port", "8080"),

			// APP密钥
			"secret": config.Env("app.secret", "123456"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("app.timezone", "Asia/Shanghai"),

			// API 域名，未设置的话所有 API URL 加 api 前缀，如 http://domain.com/api/v1/users
			"api_domain": config.Env("app.api-url", ""),
		}
	})
}
