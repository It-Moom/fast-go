/*
 * @PackageName: validators
 * @Description: 自定义验证器
 * @Author: gabbymrh
 * @Date: 2022-06-25 13:44:31
 * @LastModifiedBy: gabbymrh
 * @LastModifiedAt: 2022-06-25 13:44:31
 */

package validators

// ValidatePasswordConfirm 自定义规则，检查两次密码是否正确
func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}
