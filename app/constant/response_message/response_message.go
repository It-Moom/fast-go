/*
 * @PackageName: response_message
 * @Description: 返回消息
 * @Author: gabbymrh
 * @Date: 2021-10-02 19:01:36
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-03 16:49:59
 */

package response_message

const (

	// PARAM_ERROR 参数有误
	PARAM_ERROR = "参数有误"

	// DATA_EXIST 数据已存在
	DATA_EXIST = "数据已存在"

	// DATA_ERROR 数据有误
	DATA_ERROR = "数据有误"

	// REQUEST_SUCCESS 操作成功
	REQUEST_SUCCESS = "操作成功"

	// REQUEST_DENY 禁止访问
	REQUEST_DENY = "禁止访问"

	// PERMISSION_DENY 权限不足
	PERMISSION_DENY = "权限不足"

	// QUERY_VOID 查询为空
	QUERY_VOID = "查询为空"

	// TOKEN_EXPIRED Token无效
	TOKEN_EXPIRED = "Token无效"

	// TOO_MANY_REQUEST 请求过于频繁
	TOO_MANY_REQUEST = "请求过于频繁"

	// REQUEST_FAILS 操作失败
	REQUEST_FAILS = "操作失败"
)
