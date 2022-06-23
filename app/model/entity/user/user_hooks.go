/*
 * @PackageName: user
 * @Description: 用户实体模型钩子
 * @Author: gabbymrh
 * @Date: 2022-06-23 19:21:23
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 19:21:23
 */

package user

import (
	"fast-go/pkg/utils/secure"
	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !secure.BcryptIsHashed(userModel.Password) {
		userModel.Password = secure.BcryptHash(userModel.Password)
	}
	return
}
