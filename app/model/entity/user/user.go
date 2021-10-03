/*
 * @PackageName: user
 * @Description: 用户模型实体
 * @Author: gabbymrh
 * @Date: 2021-10-03 16:32:17
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-03 16:32:17
 */

package user

import (
	"fast-go/app/model/entity"
	"time"
)

// User 用户实体
type User struct {
	entity.BaseEntity
	Name            string    `gorm:"type:varchar(100);index;comment:用户姓名"`
	Email           string    `gorm:"type:varchar(255);unique;comment:电子邮箱"`
	PhoneNumber     string    `gorm:"type:varchar(50);unique;comment:手机号"`
	Password        string    `gorm:"type:varchar(100);comment:密码"`
	EmailVerifiedAt time.Time `gorm:"type:timestamp;comment:邮箱验证时间"`
	Status          int       `gorm:"type:int(2);not null;default:1;index;comment:状态:1=>正常,-1=>封禁"`
	Avatar          string    `gorm:"type:varchar(255);comment:头像"`
}
