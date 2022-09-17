package v1

import (
	"fast-go/app/handler/http_response"
	"fast-go/app/http/controller/api"
	httpreq "fast-go/app/http/request"
	"fast-go/app/model/entity/user"
	"fast-go/app/policy"
	"github.com/emvi/null"
	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
)

// UserController User控制器
type UserController struct {
	api.BaseAPIController
}

// GetAll 获取User列表数据
func (ctrl *UserController) GetAll(c *gin.Context) {
	userList := user.GetAll()
	http_response.RequestSuccess(c, userList)
}

// ShowOne 获取User详情数据
func (ctrl *UserController) ShowOne(c *gin.Context) {
	// 传入的参数ID转为int
	_user := user.FindByID(cast.ToInt(c.Param("id")))
	if _user.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}
	http_response.RequestSuccess(c, _user)
}

// Create 创建User数据
func (ctrl *UserController) Create(c *gin.Context) {

	// 参数校验
	request := httpreq.UserRequest{}
	if ok := httpreq.Validate(c, &request, httpreq.UserSave); !ok {
		return
	}

	userEntity := user.User{
		Email: null.NewString(request.Email, true),
		Name:  request.Name,
	}
	_user := user.Create(&userEntity)
	if _user.ID > 0 {
		http_response.RequestSuccess(c, _user)
	} else {
		http_response.RequestError(c, "创建User失败")
	}
}

// Update 更新User数据
func (ctrl *UserController) Update(c *gin.Context) {
	// 传入的参数ID转为int
	_user := user.FindByID(cast.ToInt(c.Param("id")))
	if _user.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}

	if ok := policy.CanModifyUser(c, _user); !ok {
		http_response.PermissionDeny(c, "无权修改User")
		return
	}

	// 参数校验
	request := httpreq.UserRequest{}
	if ok := httpreq.Validate(c, &request, httpreq.UserSave); !ok {
		return
	}

	userEntity := user.User{
		Name:  request.Name,
		Email: null.NewString(request.Email, true),
	}
	_user = user.Update(&userEntity)
	if _user.ID > 0 {
		http_response.RequestSuccess(c, _user)
	} else {
		http_response.RequestError(c, "修改User失败")

	}
}

// Delete 删除User数据
func (ctrl *UserController) Delete(c *gin.Context) {
	// 传入的参数ID转为int
	_user := user.FindByID(cast.ToInt(c.Param("id")))
	if _user.ID == 0 {
		http_response.QueryVoid(c, "User不存在")
		return
	}

	if ok := policy.CanModifyUser(c, _user); !ok {
		http_response.PermissionDeny(c, "无权修改User")
		return
	}

	rowsAffected := user.Delete(&_user)
	if rowsAffected > 0 {
		http_response.RequestSuccess(c, "删除User成功")
		return
	}
	http_response.RequestError(c, "删除User失败")
}
