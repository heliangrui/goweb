package validator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// Translate 翻译工具
func Translate(err error, s interface{}) string {
	var k string
	t := reflect.TypeOf(s).Elem()
	for _, errs := range err.(validator.ValidationErrors) {
		//使用反射方法获取struct种的json标签作为key --重点2
		if field, ok := t.FieldByName(errs.StructField()); ok {
			k = field.Tag.Get("errorMessage")
			if k == "" {
				k = errs.Error()
			}
		} else {
			k = errs.Error()
		}
		if k != "" {
			return k
		}
	}
	return k
}
