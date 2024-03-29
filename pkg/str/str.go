/*
 * @PackageName: str
 * @Description: 字符串处理文件
 * @Author: gabbymrh
 * @Date: 2022-06-22 23:48:18
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 23:48:18
 */

package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural 转为复数 user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular 转为单数 users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake 转为 snake_case，如 TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel 转为 CamelCase，如 topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel 转为 lowerCamelCase，如 TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
