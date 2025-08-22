package convert

import (
	"strconv"
)

func ToInt64(str string) int64 {
	converted, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return converted
}

func ToFloat64(str string) float64 {
	converted, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return converted
}

func ToBool(str string) bool {
	converted, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return converted
}

//fmt.Println(ToAny[string](1))    // "1"
//fmt.Println(ToAny[string](true)) // "true"
//fmt.Println(ToAny[string](1.1))  // "1.1"
//
//fmt.Println(ToAny[int]("1"))   	// 1
//fmt.Println(ToAny[int]("1.0")) 	// 1
//fmt.Println(ToAny[int](true))  	// 1
