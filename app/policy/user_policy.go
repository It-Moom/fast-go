/*
 * @PackageName: policy
 * @Description: 用户授权策略
 * @Author: gabbymrh
 * @Date: 2022-06-24 17:29:32
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-24 17:29:32
 */

package policy

import (
	"fast-go/app/model/entity/user"
	"fast-go/pkg/auth"
	"github.com/gin-gonic/gin"
)

// CanModifyUser 判断登录的用户是否具有修改用户数据的权限
func CanModifyUser(c *gin.Context, _topic user.User) bool {
	return auth.CurrentUID(c) == _topic.GetStringID()
}
