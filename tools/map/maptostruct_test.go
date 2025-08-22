package _map

import (
	"fmt"
	"testing"
)

type School struct {
	name string
}
type Person struct {
	Name    string
	Address string
	Phones  []string
	XInfo   string `json:"x_info"`
	School  School
}

// map 可以包含 slice 等 object 类型
// map 可以包含自定义的 struct 类型
// map 中包含的 key 如果包含 - 中划线则该字段不会被解析成功，但包含 ‘_’ 可以
func TestMapToStruct(t *testing.T) {
	m := map[string]any{
		"name":    "vey-ron",
		"address": "earth",
		"x_info":  "not-parse",
		"phones":  []string{"12345678", "87654321"}, // object type
		"school":  School{name: "cen-xi-zhong-xue"}, // nested struct
	}
	s, err := MapToStruct[string, any, Person, map[string]any](m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v", s)
}
