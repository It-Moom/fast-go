/*
 * @PackageName: secure
 * @Description: 安全工具
 * @Author: gabbymrh
 * @Date: 2021-10-03 23:24:32
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 23:24:32
 */

package secure

import (
	"fast-go/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成Bcrypt密码
func GeneratePassword(str string) string {
	// 使用Bcrypt加密密码
	password, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		logger.LogError(err)
	}
	// 返回加密字符串
	return string(password)
}

// VerifyPassword 校验密码是否相等
func VerifyPassword(password string, hashedPassword string) bool {
	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
