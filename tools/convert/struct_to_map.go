package convert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// struct to map
func ConvertToMapAny(data any) (map[string]any, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil, errors.New("data is not a struct")
	}
	t := reflect.TypeOf(data)
	var m = make(map[string]any)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			m[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return m, nil
}
func ConvertToMapString(data any) (map[string]string, error) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil, errors.New("data is not a struct")
	}
	t := reflect.TypeOf(data)
	var m = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			val, _ := ConvertToAny[string](v.Field(i).Interface())
			m[t.Field(i).Name] = val
		}
	}
	return m, nil
}
func AnyMapToMapStr(a any) (map[string]string, error) {
	var m = map[string]string{}

	switch v := a.(type) {
	case map[string]string:
		return v, nil
	case map[string]any:
		for k, val := range v {
			val, err := ConvertToAny[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]string:
		for k, val := range v {
			k, err := ConvertToAny[string](k)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case map[any]any:
		for k, val := range v {
			k, err := ConvertToAny[string](k)
			if err != nil {
				return nil, err
			}
			val, err := ConvertToAny[string](val)
			if err != nil {
				return nil, err
			}
			m[k] = val
		}
	case string:
		if err := jsonStringToObject(v, &m); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unable to convert %#v of type %T to map[string]string", a, a)
	}
	return m, nil
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v any) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
