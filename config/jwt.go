/*
 * @PackageName: config
 * @Description: JWT配置
 * @Author: gabbymrh
 * @Date: 2021-10-03 18:01:28
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 18:01:28
 */

package config

import "fast-go/pkg/config"

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			// jwt 密钥
			"secret": config.Env("jwt.secret", "fast-go"),
			// token标识名称
			"token_name": config.Env("jwt.token_name", "Authorization"),

			// jwt 有效时间:单位秒,默认1小时过期
			"ttl": config.Env("jwt.ttl", 3600),
		}
	})
}
