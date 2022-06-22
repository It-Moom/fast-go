/*
 * @PackageName: bootstrap
 * @Description: redis 引导文件
 * @Author: gabbymrh
 * @Date: 2022-06-22 17:34:13
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 17:34:13
 */

package bootstrap

import (
	"fast-go/pkg/config"
	"fast-go/pkg/redis"
	"fmt"
)

// SetupRedis 初始化 Redis
func SetupRedis() {

	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
