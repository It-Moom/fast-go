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
	"github.com/emvi/null"
	"time"
)

func MakeUser(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		entity := user.User{
			Name:            faker.Username(),
			Email:           null.NewString(faker.Email(), true),
			PhoneNumber:     null.NewString(helpers.RandomNumber(11), true),
			Password:        "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			EmailVerifiedAt: null.NewTime(time.Now(), true),
		}
		objs = append(objs, entity)
	}

	return objs
}
