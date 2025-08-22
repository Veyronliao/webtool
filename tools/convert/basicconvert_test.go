package convert

import (
	"fmt"
	"testing"
)

func TestToInt64(t *testing.T) {
	s := "123"
	V := ToInt64(s)
	if V == 123 {
		fmt.Println("TestToInt64 pass")
	}
}

func TestToFloat64(t *testing.T) {
	s := "123.5"
	f := ToFloat64(s)
	if f == 123.5 {
		fmt.Println("TestToFloat64 pass")
	}
}

func TestToBool(t *testing.T) {
	s := "true"
	b := ToBool(s)
	if b == true {
		fmt.Println("TestToBool pass")
	}
}
