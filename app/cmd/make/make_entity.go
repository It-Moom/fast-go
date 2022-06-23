/*
 * @PackageName: make
 * @Description: make entity 命令
 * @Author: gabbymrh
 * @Date: 2022-06-23 10:53:28
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 10:53:28
 */

package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CmdMakeEntity = &cobra.Command{
	Use:   "entity",
	Short: "Crate entity file, example: make entity user",
	Run:   runMakeEntity,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeEntity(cmd *cobra.Command, args []string) {

	// 格式化实体模型名称，返回一个 Entity 对象
	entity := makeEntityFromString(args[0])

	// 确保模型的目录存在，例如 `app/model/entity/user`
	dir := fmt.Sprintf("app/model/entity/%s/", entity.PackageName)
	// os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777
	os.MkdirAll(dir, os.ModePerm)

	// 替换变量
	createFileFromStub(dir+entity.PackageName+"_entity.go", "model/entity/entity", entity)
	createFileFromStub(dir+entity.PackageName+"_util.go", "model/entity/entity_util", entity)
	createFileFromStub(dir+entity.PackageName+"_hooks.go", "model/entity/entity_hooks", entity)
}
