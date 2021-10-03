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
	config.Add("jwt", config.StrMap{

		// jwt 密钥
		"jwt_secret": config.Env("JWT_SERET", "fast-go"),

		// jwt 有效时间:单位秒,默认1小时过期
		"jwt_ttl": config.Env("JWT_TTL", "3600"),
	})
}
