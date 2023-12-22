/*
 * @PackageName: bootstrap
 * @Description: 数据库引导文件
 * @Author: gabbymrh
 * @Date: 2022-06-22 17:26:53
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2022-06-22 17:26:53
 */

package bootstrap

import (
	"fast-go/pkg/config"
	"fast-go/pkg/database"
	"fast-go/pkg/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// 配置信息
	var dbEnvConfig database.DBEnvConfig
	config.ViperUnMarshal(&dbEnvConfig)
	// 数据库连接
	var dbConnection database.Connection

	// 默认数据库
	defaultDb := dbEnvConfig.Database.Default
	// 构建连接属性
	dbConnection = database.Connection{
		Host:        dbEnvConfig.Database.Connections[defaultDb].Host,
		Port:        dbEnvConfig.Database.Connections[defaultDb].Port,
		Username:    dbEnvConfig.Database.Connections[defaultDb].Username,
		Password:    dbEnvConfig.Database.Connections[defaultDb].Password,
		Database:    dbEnvConfig.Database.Connections[defaultDb].Database,
		TablePrefix: dbEnvConfig.Database.Connections[defaultDb].TablePrefix,
		Charset:     dbEnvConfig.Database.Connections[defaultDb].Charset,
	}

	// 连接数据库，并设置 GORM 的日志模式(此处采用自定义的日志模式)
	// database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.Connect(dbConnection, logger.NewGormLogger())
}
