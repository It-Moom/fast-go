package request

import (
	"fast-go/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserRequest struct {
	Name     string `valid:"name" json:"name"`
	Email    string `valid:"email" json:"email,omitempty"`
	Password string `valid:"password" json:"password,omitempty"`
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":     []string{"required", "min_cn:2", "max_cn:50", "not_exists:user,name"},
		"email":    []string{"min_cn:2", "max_cn:255", "not_exists:user,email"},
		"password": []string{"required", "min:6", "max:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字符",
			"max_cn:名称长度不能超过 50 个字符",
			"not_exists:名称已存在",
		},
		"email": []string{
			"min_cn:描述长度需至少 2 个字",
			"max_cn:描述长度不能超过 255 个字",
			"not_exists:邮箱已存在",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6 个字符",
			"max:密码长度不能超过 255 个字符",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:新的 Email 与老 Email 一致",
		},
	}

	return validate(data, rules, messages)
}
