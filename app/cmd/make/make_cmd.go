/*
 * @PackageName: make
 * @Description: make cmd 命令
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:54:14
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:54:14
 */

package make

import (
	"fast-go/pkg/console"
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, exmaple: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	entity := makeEntityFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/cmd/%s.go", entity.PackageName)

	// 从模板中创建文件（做好变量替换）
	createFileFromStub(filePath, "cmd", entity)

	// 友好提示
	console.Success("command name:" + entity.PackageName)
	console.Success("command variable name: cmd.Cmd" + entity.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
