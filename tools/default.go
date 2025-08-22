package tools

import (
	"errors"
	"reflect"
)

// SetDefaultValue
func SetDefaultValue(data any) error {
	//typeof := reflect.TypeOf(data)
	valueEle := reflect.ValueOf(data).Elem()
	if valueEle.Kind() == reflect.Ptr {
		return errors.New("must be pointer")
	}
	for i := 0; i < valueEle.NumField(); i++ {
		fieldValue := valueEle.Field(i)
		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString("")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue.SetInt(0)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldValue.SetUint(0)
		case reflect.Float32, reflect.Float64:
			fieldValue.SetFloat(0)
		case reflect.Bool:
			fieldValue.SetBool(false)
		default:
		}
	}
	return nil
}
