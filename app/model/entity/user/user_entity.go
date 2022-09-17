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
	"github.com/emvi/null"
)

// User 用户实体
type User struct {
	entity.BaseEntity
	Name            string      `gorm:"type:varchar(100);index;comment:用户姓名" json:"name,omitempty"`
	Email           null.String `gorm:"type:varchar(255);unique;comment:电子邮箱" json:"email,omitempty"`
	PhoneNumber     null.String `gorm:"type:varchar(50);unique;comment:手机号" json:"phone_number,omitempty"`
	Password        string      `gorm:"type:varchar(100);comment:密码" json:"-"`
	EmailVerifiedAt null.Time   `gorm:"comment:邮箱验证时间" json:"email_verified_at,omitempty"`
	Status          int         `gorm:"type:int(2);not null;default:1;index;comment:状态:1=>正常,-1=>封禁" json:"status,omitempty"`
	Avatar          null.String `gorm:"type:varchar(255);comment:头像" json:"avatar,omitempty"`
	entity.CommonTimestampsField
}
