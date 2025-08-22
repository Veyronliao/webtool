package safety

import (
	"fmt"
	"testing"
)

// 在 web 应用中经常会遇到数据验证问题,比较常用的的是包 validator.
// 其原理是将验证规则写在struct对应字段tag里,然后再通过反射获取struct的tag,实现数据验证.
func TestValidaStruct(t *testing.T) {
	type Person struct {
		Name string `json:"name" validate:"required"`
		Age  int64  `json:"age" validate:"required"`
	}
	p := Person{
		Name: "",
		Age:  18,
	}
	b := ValidaStruct(p)
	if b {
		fmt.Println("ValidaStruct pass")
	} else {
		fmt.Println("ValidaStruct not pass")
	}
}
