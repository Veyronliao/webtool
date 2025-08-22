package _map

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func MapToStruct[K comparable, V any, T any, M ~map[K]V](m M) (*T, error) {
	mType := reflect.TypeOf(m)
	if mType.Kind() != reflect.Map {
		return nil, errors.New("not a map")
	}
	t := new(T)
	sType := reflect.TypeOf(t)
	if sType.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	err := mapstructure.Decode(m, t)
	if err != nil {
		return t, err
	}
	return any(t).(*T), nil
}
