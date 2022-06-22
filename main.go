/*
 * @PackageName: main
 * @Description: 主入口
 * @Author: gabbymrh
 * @Date: 2021-10-02 17:35:22
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 17:35:22
 */

package main

import (
	"fast-go/app/cmd"
	"fast-go/bootstrap"
	btsConfig "fast-go/config"
	"fast-go/pkg/config"
	"fast-go/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// 初始化配置信息
	btsConfig.Initialize()
}

// 入口函数
func main() {
	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A fast web framework for Go",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdScret,
		cmd.CmdPlay,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
