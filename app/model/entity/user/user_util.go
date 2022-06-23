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
