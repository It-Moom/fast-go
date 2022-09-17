/*
 * @PackageName: tests
 * @Description: 生成器测试
 * @Author: gabbymrh
 * @Date: 2022-09-17 15:14:53
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-09-17 15:14:53
 */

package tests

import (
	"fast-go/pkg/utils/generator"
	"fmt"
	"testing"
)

// TestSnowFlakeID 测试雪花ID算法
func TestSnowFlakeID(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := generator.SnowFlakeID()
		fmt.Println(id)
	}
}
