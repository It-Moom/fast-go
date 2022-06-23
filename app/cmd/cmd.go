/*
 * @PackageName: cmd
 * @Description: cmd 命令行
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:14:28
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:14:28
 */

package cmd

import (
	"fast-go/pkg/helpers"
	"github.com/spf13/cobra"
	"os"
)

// Env 存储全局选项 --env 的值
var Env string

// RegisterGlobalFlags 注册全局选项（flag）
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "load .yml file, example: --env=testing will use .yml.testing file")
}

// RegisterDefaultCmd 注册默认命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	firstArg := helpers.FirstElement(os.Args[1:])
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}
}
