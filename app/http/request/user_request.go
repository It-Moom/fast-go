package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserRequest struct {
	Name  string `valid:"name" json:"name"`
	Email string `valid:"email" json:"email,omitempty"`
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":  []string{"required", "min_cn:2", "max_cn:50", "not_exists:user,name"},
		"email": []string{"min_cn:2", "max_cn:255", "not_exists:user,email"},
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
			"not_exists:名称已存在",
		},
	}
	return validate(data, rules, messages)
}
