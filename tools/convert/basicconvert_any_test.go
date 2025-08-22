package convert

import (
	"fmt"
	"testing"
)

func TestConvToAny(t *testing.T) {
	i := 100
	s := "100"
	i1 := true
	s1 := 1
	i2 := 123.5
	s2 := "123..5"
	v, err := ConvToAny[string](i)
	v1, err1 := ConvToAny[int](i1)
	v2, err2 := ConvToAny[string](i2)
	//v3, err3 := ConvToAny[int](i1)
	if err != nil {
		fmt.Println("TestConvToAny[string](int) not pass")
	}
	if s == v {
		fmt.Println("TestConvToAny[string](int) pass")
	}
	if err1 != nil {
		fmt.Println("TestConvToAny[int](bool) not pass")
	}
	if s1 == v1 {
		fmt.Println("TestConvToAny[int](bool) pass")
	}
	if err2 != nil {
		fmt.Println("TestConvToAny[string](float) not pass")
	}
	if s2 == v2 {
		fmt.Println("TestConvToAny[string](float) pass")
	}
}
