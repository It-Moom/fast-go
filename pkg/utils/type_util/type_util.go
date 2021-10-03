/*
 * @PackageName: type_util
 * @Description: 类型转换器
 * @Author: gabbymrh
 * @Date: 2021-10-03 16:10:06
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 16:10:06
 */

package type_util

import (
	"fast-go/pkg/logger"
	"strconv"
)

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
