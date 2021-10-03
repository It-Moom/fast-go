/*
 * @PackageName: generator
 * @Description: ID生成器
 * @Author: gabbymrh
 * @Date: 2021-10-03 21:30:32
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 21:30:32
 */

package generator

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

// SimpleUUID 简洁UUID生成器
func SimpleUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
