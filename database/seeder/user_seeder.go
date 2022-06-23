/*
 * @PackageName: seeder
 * @Description: 用户假数据种子:在通过seed 命令填充数据时可以指定单条seeder，
 * 如：go run main.go seed SeedUserTable
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:18:32
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:18:32
 */

package seeder

import (
	"fast-go/database/factory"
	"fast-go/pkg/console"
	"fast-go/pkg/logger"
	"fast-go/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedUserTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		users := factory.MakeUsers(10)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("user").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
