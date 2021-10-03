/*
 * @PackageName: bootstrap
 * @Description: 数据库初始化
 * @Author: gabbymrh
 * @Date: 2021-10-03 15:48:15
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 15:48:15
 */

package bootstrap

import (
	"fast-go/app/model/entity/user"
	"fast-go/pkg/model"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 迁移数据表
	db.AutoMigrate(&user.User{})
}
