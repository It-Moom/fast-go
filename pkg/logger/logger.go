/*
 * @PackageName: logger
 * @Description: 日志记录
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:28:58
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:28:58
 */

package logger

import "log"

// LogError 错误日志
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
