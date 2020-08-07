package util

import (
	"errors"
	"reflect"
	"strconv"
	"unsafe"
)

// []byte转string
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// string转[]byte
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func ConvertString(v interface{}) (string, error) {
	switch vv := v.(type) {
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10), nil
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10), nil
	case float64, float32:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64), nil
	case string:
		return vv, nil
	case []byte:
		return string(vv), nil
	case bool:
		return strconv.FormatBool(vv), nil
	}

	return "", errors.New("cannot convert string")
}

func MustConvertString(v interface{}) string {
	val, _ := ConvertString(v)
	return val
}

func ConvertInt64(v interface{}) (int64, error) {
	switch vv := v.(type) {
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(v).Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), nil
	case float64, float32:
		return int64(reflect.ValueOf(vv).Float()), nil
	case string:
		return strconv.ParseInt(vv, 10, 0)
	case []byte:
		return strconv.ParseInt(string(vv), 10, 0)
	case bool:
		if vv {
			return 1, nil
		}
		return 0, nil
	}

	return 0, errors.New("cannot convert int64")
}

func MustConvertInt64(v interface{}) int64 {
	val, _ := ConvertInt64(v)
	return val
}

//ConvertToInt convert some value to int
func ConvertToInt(v interface{}) (int, error) {
	int64Value, err := ConvertInt64(v)
	return int(int64Value), err
}
