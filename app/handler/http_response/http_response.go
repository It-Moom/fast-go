/*
 * @PackageName: http_response
 * @Description: Http数据返回封装
 * @Author: gabbymrh
 * @Date: 2021-10-02 18:07:32
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 18:07:32
 */

package http_response

import (
	"fast-go/app/constant/response_code"
	"fast-go/app/constant/response_message"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseBody 返回数据结构体
type responseBody struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
}

// ParamError 参数有误(默认)
func ParamError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PARAM_ERROR,
		Success: false,
		Data:    nil,
		Message: response_message.PARAM_ERROR,
		Errors:  []string{},
	})
}

// ParamErrorWithMsg 参数有误(指定message)
func ParamErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PARAM_ERROR,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// ParamErrorWithErrors 参数有误(指定errors)
func ParamErrorWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PARAM_ERROR,
		Success: false,
		Data:    nil,
		Message: response_message.PARAM_ERROR,
		Errors:  errors,
	})
}

// ParamErrorWithMsgAndErrors 参数有误(指定message和errors)
func ParamErrorWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PARAM_ERROR,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// DataExists 数据已存在(默认)
func DataExists(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_EXIST,
		Success: false,
		Data:    nil,
		Message: response_message.DATA_EXIST,
		Errors:  []string{},
	})
}

// DataExistsWithMsg 数据已存在(指定message)
func DataExistsWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_EXIST,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// DataExistsWithErrors 数据已存在(指定errors)
func DataExistsWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_EXIST,
		Success: false,
		Data:    nil,
		Message: response_message.DATA_EXIST,
		Errors:  errors,
	})
}

// DataExistsWithMsgAndErrors 数据已存在(指定message和errors)
func DataExistsWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_EXIST,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// DataError 数据错误(默认)
func DataError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_ERROR,
		Success: false,
		Data:    nil,
		Message: response_message.DATA_ERROR,
		Errors:  []string{},
	})
}

// DataErrorWithMsg 数据错误(指定message)
func DataErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_ERROR,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// DataErrorWithErrors 数据错误(指定errors)
func DataErrorWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_ERROR,
		Success: false,
		Data:    nil,
		Message: response_message.DATA_ERROR,
		Errors:  errors,
	})
}

// DataErrorWithMsgAndErrors 数据错误(指定message和errors)
func DataErrorWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.DATA_ERROR,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// RequestSuccess 操作成功(默认)
func RequestSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_SUCCESS,
		Success: true,
		Data:    data,
		Message: response_message.REQUEST_SUCCESS,
		Errors:  []string{},
	})
}

// RequestSuccessWithMsg 操作成功(指定message)
func RequestSuccessWithMsg(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_SUCCESS,
		Success: true,
		Data:    data,
		Message: msg,
		Errors:  []string{},
	})
}

// RequestDeny 禁止访问(默认)
func RequestDeny(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_DENY,
		Success: false,
		Data:    nil,
		Message: response_message.REQUEST_DENY,
		Errors:  []string{},
	})
}

// RequestDenyWithMsg 禁止访问(指定message)
func RequestDenyWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_DENY,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// RequestDenyWithErrors 禁止访问(指定errors)
func RequestDenyWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_DENY,
		Success: false,
		Data:    nil,
		Message: response_message.REQUEST_DENY,
		Errors:  errors,
	})
}

// RequestDenyWithMsgAndErrors 禁止访问(指定message和errors)
func RequestDenyWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_DENY,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// PermissionDeny 权限不足(默认)
func PermissionDeny(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PERMISSION_DENY,
		Success: false,
		Data:    nil,
		Message: response_message.PERMISSION_DENY,
		Errors:  []string{},
	})
}

// PermissionDenyWithMsg 权限不足(指定message)
func PermissionDenyWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PERMISSION_DENY,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// PermissionDenyWithErrors 权限不足(指定errors)
func PermissionDenyWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PERMISSION_DENY,
		Success: false,
		Data:    nil,
		Message: response_message.PERMISSION_DENY,
		Errors:  errors,
	})
}

// PermissionDenyWithMsgAndErrors 权限不足(指定message和errors)
func PermissionDenyWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.PERMISSION_DENY,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// QueryVoid 结果为空(默认)
func QueryVoid(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.QUERY_VOID,
		Success: false,
		Data:    nil,
		Message: response_message.PERMISSION_DENY,
		Errors:  []string{},
	})
}

// QueryVoidWithMsg 结果为空(指定message)
func QueryVoidWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.QUERY_VOID,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// QueryVoidWithErrors 结果为空(指定errors)
func QueryVoidWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.QUERY_VOID,
		Success: false,
		Data:    nil,
		Message: response_message.QUERY_VOID,
		Errors:  errors,
	})
}

// QueryVoidWithMsgAndErrors 结果为空(指定message和errors)
func QueryVoidWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.QUERY_VOID,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// TokenExpired Token无效(默认)
func TokenExpired(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOKEN_EXPIRED,
		Success: false,
		Data:    nil,
		Message: response_message.TOKEN_EXPIRED,
		Errors:  []string{},
	})
}

// TokenExpiredWithMsg Token无效(指定message)
func TokenExpiredWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOKEN_EXPIRED,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// TokenExpiredWithErrors Token无效(指定errors)
func TokenExpiredWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOKEN_EXPIRED,
		Success: false,
		Data:    nil,
		Message: response_message.TOKEN_EXPIRED,
		Errors:  errors,
	})
}

// TokenExpiredWithMsgAndErrors Token无效(指定message和errors)
func TokenExpiredWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOKEN_EXPIRED,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// TooManyRequest 请求过于频繁(默认)
func TooManyRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOO_MANY_REQUEST,
		Success: false,
		Data:    nil,
		Message: response_message.TOO_MANY_REQUEST,
		Errors:  []string{},
	})
}

// TooManyRequestWithMsg 请求过于频繁(指定message)
func TooManyRequestWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOO_MANY_REQUEST,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// TooManyRequestWithErrors 请求过于频繁(指定errors)
func TooManyRequestWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOO_MANY_REQUEST,
		Success: false,
		Data:    nil,
		Message: response_message.TOKEN_EXPIRED,
		Errors:  errors,
	})
}

// TooManyRequestWithMsgAndErrors 请求过于频繁(指定message和errors)
func TooManyRequestWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.TOO_MANY_REQUEST,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}

// RequestError 操作失败(默认)
func RequestError(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_FAILS,
		Success: false,
		Data:    nil,
		Message: response_message.REQUEST_FAILS,
		Errors:  []string{},
	})
}

// RequestErrorWithMsg 操作失败(指定message)
func RequestErrorWithMsg(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_FAILS,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  []string{},
	})
}

// RequestErrorWithErrors 操作失败(指定errors)
func RequestErrorWithErrors(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_FAILS,
		Success: false,
		Data:    nil,
		Message: response_message.TOKEN_EXPIRED,
		Errors:  errors,
	})
}

// RequestErrorWithMsgAndErrors 操作失败(指定message和errors)
func RequestErrorWithMsgAndErrors(ctx *gin.Context, msg string, errors []string) {
	ctx.JSON(http.StatusOK, responseBody{
		Code:    response_code.REQUEST_FAILS,
		Success: false,
		Data:    nil,
		Message: msg,
		Errors:  errors,
	})
}
