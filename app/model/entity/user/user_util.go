/*
 * @PackageName: user
 * @Description: 用户实体工具
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:53:57
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:53:57
 */

package user

import (
	"fast-go/pkg/app"
	"fast-go/pkg/database"
	"fast-go/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone_number = ?", phone).Count(&count)
	return count > 0
}

// FindByPhoneNumber 通过手机号来获取用户
func FindByPhoneNumber(phone string) (userEntity User) {
	database.DB.Where("phone_number = ?", phone).First(&userEntity)
	return
}

// FindByMulti 通过 手机号/Email/用户名 来获取用户
func FindByMulti(loginID string) (userEntity User) {
	database.DB.
		Where("phone_number = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userEntity)
	return
}

// FindById 通过 ID 获取用户
func FindById(idstr string) (userEntity User) {
	database.DB.Where("id", idstr).First(&userEntity)
	return
}

// FindByEmail 通过 Email 来获取用户
func FindByEmail(email string) (userEntity User) {
	database.DB.Where("email = ?", email).First(&userEntity)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paginate paginator.Paging) {
	paginate = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.URL(database.TableName(&User{})),
		perPage,
	)
	return
}
