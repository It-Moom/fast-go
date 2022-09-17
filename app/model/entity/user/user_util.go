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
	"fast-go/pkg/utils/secure"
	"github.com/gin-gonic/gin"
)

// GetAll 查询所有用户
func GetAll() (users []User) {
	database.DB.Find(&users)
	return users
}

// FindByID 根据ID查询用户,找不到则返回空
func FindByID(id int) (user User) {
	result := database.DB.Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 || result.Error != nil {
		return User{}
	}
	return user
}

// Create 创建用户，返回创建的用户
func Create(userEntity *User) (user User) {
	database.DB.Create(&userEntity)
	return *userEntity
}

// Update 更新用户,返回保存的用户
func Update(userEntity *User) (user User) {
	database.DB.Save(&userEntity)
	return *userEntity
}

// Delete 删除用户
func Delete(userEntity *User) (rowsAffected int64) {
	result := database.DB.Delete(&userEntity)
	return result.RowsAffected
}

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

// FindByEmail 通过 Email 来获取用户
func FindByEmail(email string) (userEntity User) {
	database.DB.Where("email = ?", email).First(&userEntity)
	return
}

// ComparePassword 密码是否正确
func (userEntity *User) ComparePassword(_password string) bool {
	return secure.BcryptCheck(_password, userEntity.Password)
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
