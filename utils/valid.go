package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 返回结构体中的 msg 参数
func GetValidMsg(err error, obj any) string {
	// 使用的时候，需要传 obj 的指针
	getObj := reflect.TypeOf(obj)

	// 将 err 接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据校验字段，获取结构体的具体字段
			if f, exists := getObj.Elem().FieldByName(e.Field()); exists {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
