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
			"host":     config.Env("database.host", "127.0.0.1"),
			"port":     config.Env("database.port", "3306"),
			"database": config.Env("database.database", "fast_go"),
			"username": config.Env("database.username", ""),
			"password": config.Env("database.password", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Env("database.max-idle-connections", 100),
			"max_open_connections": config.Env("database.max-open-connections", 25),
			"max_life_seconds":     config.Env("databse.max-life-seconds", 5*60),
		},
	})
}
