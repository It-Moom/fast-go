/*
 * @PackageName: make
 * @Description: make migration 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 15:01:33
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 15:01:33
 */

package make

import (
	"fast-go/pkg/app"
	"fast-go/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeMigration = &cobra.Command{
	Use:   "migration",
	Short: "Create a migration file, example: make migration add_user_table",
	Run:   runMakeMigration,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeMigration(cmd *cobra.Command, args []string) {

	// 日期格式化
	timeStr := app.TimenowInTimezone().Format("2006_01_02_150405")

	// 格式化模型名称，返回一个 Entity 对象
	entity := makeEntityFromString(args[0])
	fileName := timeStr + "_" + entity.PackageName
	filePath := fmt.Sprintf("database/migrations/%s.go", fileName)
	createFileFromStub(filePath, "migration", entity, map[string]string{"{{FileName}}": fileName})
	console.Success("Migration file created，after modify it, use `migrate up` to migrate database.")
}
