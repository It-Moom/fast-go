/*
 * @PackageName: cmd
 * @Description: migrate 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 14:54:18
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 14:54:18
 */

package cmd

import (
	"fast-go/database/migrations"
	"fast-go/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
	// 所有 migrate 下的子命令都会执行以下代码
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下的所有迁移文件
	migrations.Initialize()
	// 初始化 migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}
