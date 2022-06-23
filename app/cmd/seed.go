/*
 * @PackageName: cmd
 * @Description: seed 命令:在通过seed 命令填充数据时可以指定单条seeder，
 * 如：go run main.go seed SeedUserTable
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:20:58
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:20:58
 */

package cmd

import (
	sd "fast-go/database/seeder"
	"fast-go/pkg/console"
	"fast-go/pkg/seed"
	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Insert fake data to the database",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1), // 只允许 1 个参数
}

func runSeeders(cmd *cobra.Command, args []string) {
	sd.Initialize()
	if len(args) > 0 {
		// 有传参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 默认运行全部迁移
		seed.RunAll()
		console.Success("Done seeding.")
	}
}
