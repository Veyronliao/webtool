package safety

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var Validater *validator.Validate

func init() {
	Validater = validator.New()
}

func BasicValida() {

}

// 使用validator校验struct
func ValidaStruct(data any) bool {
	//判断是否是struct
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Struct {
		return false
	}
	err := Validater.Struct(data)
	if err != nil {
		return false
	} else {
		return true
	}
}
