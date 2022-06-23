/*
 * @PackageName: seeder
 * @Description: 种子数据初始化配置
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:19:28
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:19:28
 */

package seeder

import "fast-go/pkg/seed"

func Initialize() {

	// 触发加载本目录下其他文件中的 init 方法

	// 指定优先于同目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeedUserTable",
	})
}
