/*
 * @PackageName: tests
 * @Description: JWT工具单元测试
 * @Author: gabbymrh
 * @Date: 2021-10-03 19:00:55
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2022-06-24 19:39:37
 */

package tests

import (
	"fast-go/pkg/utils/jwt"
	"testing"
)

func TestJwtCreate(t *testing.T) {
	token := jwt.NewJWT().IssueToken("1")
	t.Log(token)
}
