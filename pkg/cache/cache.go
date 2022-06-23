/*
 * @PackageName: cache
 * @Description: 缓存工具
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:31:09
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:31:09
 */

package cache

import (
	"encoding/json"
	"fast-go/pkg/logger"
	"github.com/spf13/cast"
	"sync"
	"time"
)

// CacheService 缓存服务
type CacheService struct {
	Store Store
}

var once sync.Once
var Cache *CacheService

// InitWithCacheStore 初始化缓存服务
func InitWithCacheStore(store Store) {
	once.Do(func() {
		Cache = &CacheService{
			Store: store,
		}
	})
}

// Set 设置缓存
func Set(key string, obj interface{}, expireTime time.Duration) {
	b, err := json.Marshal(&obj)
	logger.LogIf(err)
	Cache.Store.Set(key, string(b), expireTime)
}

// Get 获取缓存
func Get(key string) interface{} {
	stringValue := Cache.Store.Get(key)
	var wanted interface{}
	err := json.Unmarshal([]byte(stringValue), &wanted)
	logger.LogIf(err)
	return wanted
}

// Has 检查缓存是否存在
func Has(key string) bool {
	return Cache.Store.Has(key)
}

// GetObject 应该传地址，用法如下:
//     entity := user.User{}
//     cache.GetObject("key", &entity)
func GetObject(key string, wanted interface{}) {
	val := Cache.Store.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogIf(err)
	}
}

// GetString 获取缓存字符串
func GetString(key string) string {
	return cast.ToString(Get(key))
}

// GetBool 获取缓存布尔值
func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

// GetInt 获取缓存整数
func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

// GetInt32 获取缓存整数
func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

// GetInt64 获取缓存整数
func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

// GetUint 获取缓存整数
func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

// GetUint32 获取缓存整数
func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

// GetUint64 获取缓存整数
func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

// GetFloat64 获取缓存浮点数
func GetFloat64(key string) float64 {
	return cast.ToFloat64(Get(key))
}

// GetTime 获取缓存时间
func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

// GetDuration 获取缓存时间
func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

// GetIntSlice 获取缓存整数列表
func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

// GetStringSlice 获取缓存字符串列表
func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

// GetStringMap 获取缓存字符串map
func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}

// GetStringMapString 获取缓存字符串map
func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

// GetStringMapStringSlice 获取缓存字符串map列表
func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}

// Forget 删除缓存
func Forget(key string) {
	Cache.Store.Forget(key)
}

// Forever 持久化缓存
func Forever(key string, value string) {
	Cache.Store.Set(key, value, 0)
}

// Flush 删除所有缓存
func Flush() {
	Cache.Store.Flush()
}

// Increment 缓存递增
func Increment(parameters ...interface{}) {
	Cache.Store.Increment(parameters...)
}

// Decrement 缓存递减
func Decrement(parameters ...interface{}) {
	Cache.Store.Decrement(parameters...)
}

// IsAlive 检查缓存是否有效
func IsAlive() error {
	return Cache.Store.IsAlive()
}
