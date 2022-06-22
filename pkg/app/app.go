/*
 * @PackageName: app
 * @Description: app 应用
 * @Author: gabbymrh
 * @Date: 2022-06-22 17:26:53
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2022-06-22 17:26:53
 */

package app

import "fast-go/pkg/config"

// IsLocal 是否本地环境
func IsLocal() bool {
	return config.Get("app.env") == "local"
}

// IsProduction 是否生产环境
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

// IsTest 是否测试环境
func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
