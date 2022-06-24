/*
 * @PackageName: auth
 * @Description: 用户认证
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:16:26
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:16:26
 */

package auth

import (
	"errors"
	"fast-go/app/model/entity/user"
	"fast-go/pkg/logger"
	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userEntity := user.FindByMulti(email)
	if userEntity.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userEntity.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userEntity, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userEntity := user.FindByPhoneNumber(phone)
	if userEntity.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userEntity, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userEntity, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userEntity
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
