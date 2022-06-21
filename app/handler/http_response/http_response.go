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
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// ResponseBody 返回数据结构体
type responseBody struct {
	Code    int         `json:"code"`    // 返回code
	Success bool        `json:"success"` // 成功状态:true/false
	Data    interface{} `json:"data"`    // 数据体
	Message string      `json:"message"` // 提示消息
	Errors  []string    `json:"errors"`  // 错误列表(若有)
}

// ParamError 参数有误
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func ParamError(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.PARAM_ERROR, response_message.PARAM_ERROR, args...)
}

// DataExists 数据已存在
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func DataExists(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.DATA_EXIST, response_message.DATA_EXIST, args...)
}

// DataError 数据错误
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func DataError(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.DATA_ERROR, response_message.DATA_ERROR, args...)
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

// RequestDeny 禁止访问
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func RequestDeny(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.REQUEST_DENY, response_message.REQUEST_DENY, args...)
}

// PermissionDeny 权限不足
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func PermissionDeny(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.PERMISSION_DENY, response_message.PERMISSION_DENY, args...)
}

// QueryVoid 结果为空
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func QueryVoid(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.QUERY_VOID, response_message.QUERY_VOID, args...)
}

// TokenExpired Token无效
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func TokenExpired(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.TOKEN_EXPIRED, response_message.TOKEN_EXPIRED, args...)
}

// TooManyRequest 请求过于频繁
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func TooManyRequest(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.TOO_MANY_REQUEST, response_message.TOO_MANY_REQUEST, args...)
}

// RequestError 操作失败
// (函数接收1-2个参数，只接受string或[]string类型错误参数，如果为2个参数,则第1个参数必须是string，第二个参数必须是[]string)
func RequestError(ctx *gin.Context, args ...interface{}) {
	responseResult(ctx, response_code.REQUEST_FAILS, response_message.REQUEST_FAILS, args...)
}

// responseResult 响应结果
func responseResult(ctx *gin.Context, resCode int, resMsg string, args ...interface{}) {
	// 取出args长度
	argsLen := len(args)

	// 通过switch判断参数类型
	switch {

	// 如果传入参数数量等于1，则判断参数类型
	case argsLen == 1:
		// 如果参数类型为string，则直接返回message
		if reflect.TypeOf(args[0]).String() == "string" {
			ctx.JSON(http.StatusOK, responseBody{
				Code:    resCode,
				Success: false,
				Data:    nil,
				Message: args[0].(string),
				Errors:  []string{},
			})
		}
		// 如果参数类型为[]string，则直接返回errors
		if reflect.TypeOf(args[0]).String() == "[]string" {
			ctx.JSON(http.StatusOK, responseBody{
				Code:    resCode,
				Success: false,
				Data:    nil,
				Message: resMsg,
				Errors:  args[0].([]string),
			})
		}

	// 如果传入参数数量等于2，则判断参数类型
	case argsLen == 2:
		// 如果第一个参数类型为string，则直接返回message,errors
		if reflect.TypeOf(args[0]).String() == "string" {
			ctx.JSON(http.StatusOK, responseBody{
				Code:    resCode,
				Success: false,
				Data:    nil,
				Message: args[0].(string),
				Errors:  args[1].([]string),
			})
		}

	// 返回默认消息
	default:
		ctx.JSON(http.StatusOK, responseBody{
			Code:    resCode,
			Success: false,
			Data:    nil,
			Message: resMsg,
			Errors:  []string{},
		})
	}
}
