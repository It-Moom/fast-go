package v1

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/http/controller/api"
	httpreq "fast-go/app/http/request"
	"fast-go/app/model/entity/user"
	"fast-go/app/policy"

	"github.com/gin-gonic/gin"
)

// UserController User控制器
type UserController struct {
	api.BaseAPIController
}

// ListAll 获取User列表数据
func (ctrl *UserController) ListAll(c *gin.Context) {
	users := user.All()
	http_response.RequestSuccess(c, users)
}

// ShowOne 获取User详情数据
func (ctrl *UserController) ShowOne(c *gin.Context) {
	userEntity := user.FindById(c.Param("id"))
	if userEntity.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}
	http_response.RequestSuccess(c, userEntity)
}

// Store 创建User数据
func (ctrl *UserController) Store(c *gin.Context) {
	//
	//// 参数校验
	//request := httpreq.UserRequest{}
	//if ok := httpreq.Validate(c, &request, httpreq.UserSave); !ok {
	//	return
	//}
	//
	//userEntity := user.User{
	//	Email: request.Email,
	//	Name:  request.Name,
	//}
	//userEntity.Create()
	//if userEntity.ID > 0 {
	//	http_response.RequestSuccess(c, userEntity)
	//} else {
	//	http_response.RequestError(c, "创建User失败")
	//}
}

// Update 更新User数据
func (ctrl *UserController) Update(c *gin.Context) {

	userEntity := user.FindById(c.Param("id"))
	if userEntity.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}

	if ok := policy.CanModifyUser(c, userEntity); !ok {
		http_response.PermissionDeny(c, "无权修改User")
		return
	}

	// 参数校验
	request := httpreq.UserRequest{}
	if ok := httpreq.Validate(c, &request, httpreq.UserSave); !ok {
		return
	}

	userEntity.Name = request.Name
	userEntity.Email = request.Email
	rowsAffected := userEntity.Save()
	if rowsAffected > 0 {
		http_response.RequestSuccess(c, userEntity)
	} else {
		http_response.RequestError(c, "修改User失败")

	}
}

// Delete 删除User数据
func (ctrl *UserController) Delete(c *gin.Context) {

	userEntity := user.FindById(c.Param("id"))
	if userEntity.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}

	if ok := policy.CanModifyUser(c, userEntity); !ok {
		http_response.PermissionDeny(c, "无权修改User")
		return
	}

	rowsAffected := userEntity.Delete()
	if rowsAffected > 0 {
		http_response.RequestSuccess(c, "删除User成功")
		return
	}
	http_response.RequestError(c, "删除User失败")
}
