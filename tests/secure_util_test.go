/*
 * @PackageName: tests
 * @Description: 安全工具单元测试
 * @Author: gabbymrh
 * @Date: 2021-10-03 23:28:51
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 23:28:51
 */

package tests

import (
	"fast-go/pkg/utils/secure"
	"testing"
)

// 测试生成密码
func TestGeneratePassword(t *testing.T) {
	password := secure.BcryptHash("123456")
	t.Log(password) // 如:$2a$10$caUw3B4E1JmevafFax6UyuI7OqRW4AJlWxlhvCptlTAG0fvqLDS6m
}

// 测试校验密码
func TestVerifyPassword(t *testing.T) {
	hashedPassword := "$2a$10$caUw3B4E1JmevafFax6UyuI7OqRW4AJlWxlhvCptlTAG0fvqLDS6m"
	passwordVerify := secure.BcryptCheck("123456", hashedPassword)
	t.Log(passwordVerify) // 如：true
}
