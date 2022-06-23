/*
 * @PackageName: make
 * @Description: make 命令
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:50:14
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:50:14
 */

package make

import (
	"embed"
	"fast-go/pkg/console"
	"fast-go/pkg/file"
	"fast-go/pkg/str"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"strings"
)

// Entity 参数解释
// 单个词，用户命令传参，以 User 模型为例：
//     - user
//  - User
// 整理好的数据：
// {
//     "TableName": "user",
//     "StructName": "User",
//     "VariableName": "user",
//     "PackageName": "user"
// }
// -
// 两个词或者以上，用户命令传参，以 TopicComment 模型为例：
//     - topic_comment
//  - TopicComment
// 整理好的数据：
// {
//     "TableName": "topic_comment",
//     "StructName": "TopicComment",
//     "VariableName": "topicComment",
//     "PackageName": "topic_comment"
// }
type Entity struct {
	TableName    string
	StructName   string
	VariableName string
	PackageName  string
}

// stubsFS 方便我们后面打包这些 .stub 为后缀名的文件

// 注意，下面注释不能加空格隔开，否则会导致命令行解析错误，该命令是告诉程序寻找stubs目录下的文件
//go:embed stubs
var stubsFS embed.FS

// CmdMake 说明 cobra 命令
var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册 make 的子命令
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeEntity,
		CmdMakeAPIController,
	)
}

// makeEntityFromString 格式化用户输入的内容
func makeEntityFromString(name string) Entity {
	entity := Entity{}
	entity.StructName = str.Singular(strcase.ToCamel(name))
	entity.TableName = str.Snake(entity.StructName)
	entity.VariableName = str.LowerCamel(entity.StructName)
	entity.PackageName = str.Snake(entity.StructName)
	return entity
}

// createFileFromStub 读取 stub 文件并进行变量替换
// 最后一个选项可选，如若传参，应传 map[string]string 类型，作为附加的变量搜索替换
func createFileFromStub(filePath string, stubName string, entity Entity, variables ...interface{}) {

	// 实现最后一个参数可选
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	// 目标文件已存在
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists!")
	}

	// 读取 stub 模板文件
	entityData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}
	entityStub := string(entityData)

	// 添加默认的替换变量
	replaces["{{VariableName}}"] = entity.VariableName
	replaces["{{StructName}}"] = entity.StructName
	replaces["{{PackageName}}"] = entity.PackageName
	replaces["{{TableName}}"] = entity.TableName

	// 对模板内容做变量替换
	for search, replace := range replaces {
		entityStub = strings.ReplaceAll(entityStub, search, replace)
	}

	// 存储到目标文件中
	err = file.Put([]byte(entityStub), filePath)
	if err != nil {
		console.Exit(err.Error())
	}

	// 提示成功
	console.Success(fmt.Sprintf("[%s] created.", filePath))
}
