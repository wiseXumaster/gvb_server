package utils

import "github.com/go-playground/validator/v10"

// 自定义验证函数
func IsBool(fl validator.FieldLevel) bool {
	// 对 bool 值进行验证，true 或 false 均为有效值
	value := fl.Field().Bool()
	return value == true || value == false
}
