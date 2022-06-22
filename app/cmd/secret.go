/*
 * @PackageName: cmd
 * @Description: app secret 命令
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:20:13
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:20:13
 */

package cmd

import (
	"fast-go/pkg/console"
	"fast-go/pkg/helpers"
	"github.com/spf13/cobra"
)

var CmdScret = &cobra.Command{
	Use:   "secret",
	Short: "Generate App Secret, will print the generated Secret",
	Run:   runSecretGenerate,
	Args:  cobra.NoArgs, // 不允许传参
}

func runSecretGenerate(cmd *cobra.Command, args []string) {
	console.Success("---")
	console.Success("App Secret:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .yml file to change the app.secret option")
}
