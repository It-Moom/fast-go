/*
 * @PackageName: make
 * @Description: make factory 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:30:37
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:30:37
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeFactory = &cobra.Command{
	Use:   "factory",
	Short: "Create model's factory file, exmaple: make factory user",
	Run:   runMakeFactory,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeFactory(cmd *cobra.Command, args []string) {

	// 格式化模型名称，返回一个 Entity 对象
	entity := makeEntityFromString(args[0])

	// 拼接目标文件路径
	filePath := fmt.Sprintf("database/factory/%s_factory.go", entity.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "factory", entity)
}
