/*
 * @PackageName: make
 * @Description: make service 命令
 * @Author: gabbymrh
 * @Date: 2022-06-24 09:25:44
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-24 09:25:44
 */

package make

import (
	"fast-go/app/constant/system"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// CmdMakeService 创建模型的 service 命令
var CmdMakeService = &cobra.Command{
	Use:   "service",
	Short: "Crate service intf and impl file, example: make service user",
	Run:   runMakeService,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeService(cmd *cobra.Command, args []string) {

	// 格式化实体模型名称，返回一个 Entity 对象
	entity := makeEntityFromString(args[0])

	// 确保intf目录存在
	intfDir := fmt.Sprintf("app/service/%s/", system.SERVICE_INTF_DIR_NAME)
	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll(intfDir, os.ModePerm)

	// 确保impl的目录存在
	implDir := fmt.Sprintf("app/service/%s/", system.SERVICE_IMPL_DIR_NAME)
	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll(implDir, os.ModePerm)

	// 替换变量
	createFileFromStub(intfDir+entity.PackageName+"_intf.go", "service/intf", entity)
	createFileFromStub(implDir+entity.PackageName+"_impl.go", "service/impl", entity)
}
