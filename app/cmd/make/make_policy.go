/*
 * @PackageName: make
 * @Description: make policy 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:13:14
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:13:14
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakePolicy(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Model 对象
	entity := makeEntityFromString(args[0])

	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll("app/policies", os.ModePerm)

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/policy/%s_policy.go", entity.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "policy", entity)
}
