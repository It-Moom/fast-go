/*
 * @PackageName: database
 * @Description: 数据库处理
 * @Author: gabbymrh
 * @Date: 2022-06-22 17:26:53
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2023-12-18 17:26:53
 */

package database

import (
	"database/sql"
	"errors"
	"fast-go/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// DBEnvConfig 数据库环境配置
type DBEnvConfig struct {
	// 数据库配置
	Database struct {
		// 默认数据库
		Default string
		// 数据库连接池
		Connections map[string]Connection
	}
}

// Connection 数据库连接配置
type Connection struct {
	// 驱动
	Driver string
	// 主机
	Host string
	// 端口
	Port int
	// 数据库名
	Database string
	// 用户名
	Username string
	// 密码
	Password string
	// 字符集
	Charset string
	// 最大连接数
	MaxOpenConn int
	// 最大空闲连接数
	MaxIdleConn int
	// 连接过期时间
	MaxLifetime int
}

// Connect 连接数据库
func Connect(connection Connection, _logger gormlogger.Interface) {
	// 默认数据库驱动
	if connection.Driver == "" {
		connection.Driver = "mysql"
	}
	// 遍历数据库驱动
	switch connection.Driver {
	case "mysql":
		// 连接 mysql
		mySqlConnect(connection, _logger)
	case "sqlite":
		// 连接 sqlite
		db := connection.Database
		sqlite.Open(db)
	default:
		// 未知驱动
		panic(errors.New("Unknown driver"))
	}
}

// mySqlConnect 连接 mysql
func mySqlConnect(connection Connection, _logger gormlogger.Interface) {
	// 设置默认字符集
	if connection.Charset == "" {
		connection.Charset = "utf8mb4"
	}
	// 设置默认端口
	if connection.Port == 0 {
		connection.Port = 3306
	}
	// 设置默认最大连接数
	if connection.MaxOpenConn == 0 {
		connection.MaxOpenConn = 100
	}
	// 设置默认最大空闲连接数
	if connection.MaxIdleConn == 0 {
		connection.MaxIdleConn = 10
	}
	// 设置默认连接过期时间
	if connection.MaxLifetime == 0 {
		connection.MaxLifetime = 60
	}
	// 构建DSN信息
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		connection.Username,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Database,
		connection.Charset,
	)
	dbConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
		// 关闭表名复数形式
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println("获取底层SQL DB出错" + err.Error())
	}
	// 设置最大连接数
	SQLDB.SetMaxOpenConns(connection.MaxOpenConn)
	// 设置最大空闲连接数
	SQLDB.SetMaxIdleConns(connection.MaxIdleConn)
	// 设置连接过期时间
	SQLDB.SetConnMaxLifetime(time.Duration(connection.MaxLifetime) * time.Second)
}

// SwitchDB 切换数据源
func SwitchDB(connName string) *gorm.DB {
	// 关闭当前数据库连接
	err := SQLDB.Close()
	if err != nil {
		fmt.Println("切换数据源出错" + err.Error())
	}

	// 重新连接数据库,获取数据库配置
	var dbEnvConfig DBEnvConfig
	config.ViperUnMarshal(&dbEnvConfig)
	conn := dbEnvConfig.Database.Connections[connName]
	if conn.Database == "" {
		panic(errors.New("Unknown Connection"))
	}
	// 数据库连接
	Connect(conn, nil)
	return DB
}

// CurrentDatabase 获取当前数据库名
func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

// DeleteAllTables 删除所有数据表
func DeleteAllTables() error {
	// 判断当前数据库类型
	switch DB.Dialector.Name() {
	case "mysql":
		return deleteMySQLTables()
	case "sqlite":
		return deleteAllSqliteTables()
	default:
		return errors.New("Unknown database type")
	}
}

func deleteAllSqliteTables() error {
	tables := []string{}

	// 读取所有数据表
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	tables := []string{}

	// 读取所有数据表
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}

// TableName 获取表名
func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
