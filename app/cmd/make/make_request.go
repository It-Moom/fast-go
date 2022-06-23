/*
 * @PackageName: make
 * @Description: make request 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 12:26:19
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 12:26:19
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeRequest = &cobra.Command{
	Use:   "request",
	Short: "Create request file, example make request user",
	Run:   runMakeRequest,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeRequest(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Entity 对象
	entity := makeEntityFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/http/request/%s_request.go", entity.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "request", entity)
}
