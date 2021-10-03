/*
 * @PackageName: config
 * @Description: 数据库配置
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:30:45
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:30:45
 */

package config

import "fast-go/pkg/config"

func init() {

	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{

			// 数据库连接信息
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", "fast_go"),
			"username": config.Env("DB_USERNAME", ""),
			"password": config.Env("DB_PASSWORD", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}
