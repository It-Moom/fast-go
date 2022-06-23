/*
 * @PackageName: bootstrap
 * @Description: 缓存引导程序
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:38:26
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:38:26
 */

package bootstrap

import (
	"fast-go/pkg/cache"
	"fast-go/pkg/config"
	"fmt"
)

// SetupCache 缓存
func SetupCache() {

	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
	)

	cache.InitWithCacheStore(rds)
}
