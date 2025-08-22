package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

func indirect(a any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
func ConvToString(data any) (string, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		return val, nil
	case []byte:
		return string(val), nil
	case bool:
		return strconv.FormatBool(val), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(val, 10), nil
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), nil
	case json.Number:
		return val.String(), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", data, val)
	}
}
func ConvToInt(data any) (int, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, err
		} else {
			return int(v), nil
		}
	case []byte:
		v, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return 0, err
		} else {
			return int(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return int(val.(int)), nil
	case nil:
		return 0, nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return int(v), nil
		}
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", data, val)
	}
}
func ConvToInt64(data any) (int64, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		return strconv.ParseInt(val, 10, 64)
	case []byte:
		return strconv.ParseInt(string(val), 10, 64)
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return int64(val.(int)), nil
	case nil:
		return 0, nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return int64(v), nil
		}
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", data, val)
	}
}
func ConvToInt32(data any) (int32, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0, err
		} else {
			return int32(v), nil
		}
	case []byte:
		v, err := strconv.ParseInt(string(val), 10, 32)
		if err != nil {
			return 0, err
		} else {
			return int32(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return int32(val.(int)), nil
	case nil:
		return 0, nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return int32(v), nil
		}
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", data, val)
	}
}
func ConvToInt16(data any) (int16, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseInt(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return int16(v), nil
		}
	case []byte:
		v, err := strconv.ParseInt(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return int16(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return int16(val.(int)), nil
	case nil:
		return 0, nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return int16(v), nil
		}
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", data, val)
	}

}
func ConvToInt8(data any) (int8, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		i, err := strconv.ParseInt(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return int8(i), nil
		}
	case []byte:
		i, err := strconv.ParseInt(string(val), 10, 8)
		if err != nil {
			return 0, err
		} else {
			return int8(i), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return int8(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		return int8(v), err
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", data, val)
	}

}
func ConvToFloat64(data any) (float64, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		return strconv.ParseFloat(val, 64)
	case []byte:
		return strconv.ParseFloat(string(val), 64)
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return float64(val.(int)), nil
	case json.Number:
		v, err := val.Float64()
		if err != nil {
			return 0, err
		} else {
			return v, nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", data, val)
	}

}
func ConvToFloat32(data any) (float32, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return 0, err
		} else {
			return float32(v), nil
		}
	case []byte:
		v, err := strconv.ParseFloat(string(val), 32)
		if err != nil {
			return 0, err
		} else {
			return float32(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return float32(val.(int)), nil
	case json.Number:
		v, err := val.Float64()
		if err != nil {
			return 0, err
		} else {
			return float32(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float32", data, val)
	}
}
func ConvToBool(data any) (bool, error) {
	data = indirect(data)
	switch val := data.(type) {
	case bool:
		if val {
			return true, nil
		} else {
			return false, nil
		}
	case nil:
		return false, nil
	case string:
		i, err := strconv.ParseBool(val)
		if err != nil {
			return false, err
		} else {
			return i, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return !reflect.ValueOf(val).IsZero(), nil
	case time.Duration:
		return val != 0, nil
	case json.Number:
		v, err := val.Float64()
		return v != 0, err
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", data, val)
	}
}
func ConvToUint(data any) (uint, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint(v), nil
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return uint(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return uint(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", data, val)

	}
}
func ConvToUint64(data any) (uint64, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint64(v), nil
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint64(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return uint64(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return uint64(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", data, val)
	}
}
func ConvToUint32(data any) (uint32, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint32(v), nil
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint32(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return uint32(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return uint32(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", data, val)

	}
}
func ConvToUint16(data any) (uint16, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint16(v), nil
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint16(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return uint16(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return uint16(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", data, val)

	}
}
func ConvToUint8(data any) (uint8, error) {
	data = indirect(data)
	switch val := data.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint8(v), nil
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err != nil {
			return 0, err
		} else {
			return uint8(v), nil
		}
	case bool:
		if val {
			return 1, nil
		} else {
			return 0, nil
		}
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, uintptr, complex64, complex128:
		return uint8(val.(int)), nil
	case json.Number:
		v, err := val.Int64()
		if err != nil {
			return 0, err
		} else {
			return uint8(v), nil
		}
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", data, val)

	}
}
func ConvToAny[T any](data any) (T, error) {
	var t T //声明T的变量，将转换成any类，再读取t的type属性
	switch any(t).(type) {
	case string:
		v, err := ConvToString(any(t))
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case bool:
		v, err := ConvToBool(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T) //先将v转换成any类型，再断言成T
		}
	case float64:
		v, err := ConvToFloat64(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case float32:
		v, err := ConvToFloat32(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case int:
		v, err := ConvToInt(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case int64:
		v, err := ConvToInt64(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case int32:
		v, err := ConvToInt32(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case int16:
		v, err := ConvToInt16(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case int8:
		v, err := ConvToInt8(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case uint:
		v, err := ConvToUint(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case uint64:
		v, err := ConvToUint64(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case uint32:
		v, err := ConvToUint32(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case uint16:
		v, err := ConvToUint16(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	case uint8:
		v, err := ConvToUint8(data)
		if err != nil {
			return t, err
		} else {
			t = any(v).(T)
		}
	default:
		return t, fmt.Errorf("the type %T isn't supported", t)
	}
	return t, nil
}
