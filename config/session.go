/*
 * @PackageName: config
 * @Description: 会话配置
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:30:17
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:30:17
 */

package config

import "fast-go/pkg/config"

func init() {
	config.Add("session", config.StrMap{

		// 目前只支持 Cookie
		"default": config.Env("SESSION_DRIVER", "cookie"),

		// 会话的 Cookie 名称
		"session_name": config.Env("SESSION_NAME", "fast-go-session"),
	})
}
