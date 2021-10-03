/*
 * @PackageName: response_code
 * @Description: 返回code常量
 * @Author: gabbymrh
 * @Date: 2021-10-02 18:32:55
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-02 18:32:55
 */

package response_code

const (

	// PARAM_ERROR 参数有误
	PARAM_ERROR = 10001

	// DATA_EXIST 数据已存在
	DATA_EXIST = 10002

	// DATA_ERROR 数据有误
	DATA_ERROR = 10003

	// REQUEST_SUCCESS 操作成功
	REQUEST_SUCCESS = 20000

	// REQUEST_DENY 禁止访问
	REQUEST_DENY = 40001

	// PERMISSION_DENY 权限不足
	PERMISSION_DENY = 40003

	// QUERY_VOID 查询为空
	QUERY_VOID = 40004

	// TOKEN_EXPIRED Token无效
	TOKEN_EXPIRED = 40005

	// TOO_MANY_REQUEST 请求过于频繁
	TOO_MANY_REQUEST = 40029

	// REQUEST_FAILS 操作失败
	REQUEST_FAILS = 50000
)
