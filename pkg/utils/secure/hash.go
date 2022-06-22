/*
 * @PackageName: secure
 * @Description: hash 工具
 * @Author: gabbymrh
 * @Date: 2022-06-22 18:14:29
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-22 18:14:29
 */

package secure

import (
	"fast-go/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对字符串进行加密
func BcryptHash(plainText string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文字符串和密文哈希值
func BcryptCheck(plainText, cipherText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(cipherText), []byte(plainText))
	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}
