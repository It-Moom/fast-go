/*
 * @PackageName: factory
 * @Description: 用户数据工厂
 * @Author: gabbymrh
 * @Date: 2022-06-23 18:10:12
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-23 18:10:12
 */

package factory

import (
	"fast-go/app/model/entity/user"
	"fast-go/pkg/helpers"
	"github.com/bxcodec/faker/v3"
	"time"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:            faker.Username(),
			Email:           faker.Email(),
			PhoneNumber:     helpers.RandomNumber(11),
			Password:        "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			EmailVerifiedAt: time.Now(),
		}
		objs = append(objs, model)
	}

	return objs
}
