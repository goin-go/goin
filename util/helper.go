package util

import (
	"path"
	"reflect"
	"strconv"
)

// JoinPaths 拼接路径
func JoinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := LastChar(relativePath) == '/' && LastChar(finalPath) != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}

// FirstChar 字符串首字符
func FirstChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[0]
}

// LastChar 字符串最后一个字符
func LastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string can't be 0")
	}
	return str[len(str)-1]
}

// InterfaceToString 接口类型转字符串
func InterfaceToString(val interface{}) (string, bool) {
	v := reflect.ValueOf(val)
	ok := true
	var value string

	switch v.Kind() {
	case reflect.Bool:
		value = strconv.FormatBool(v.Bool())
		break
	case reflect.String:
		value = v.String()
		break
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value = strconv.FormatInt(v.Int(), 10)
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value = strconv.FormatUint(v.Uint(), 10)
		break
	case reflect.Float64:
		value = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		break
	case reflect.Float32:
		value = strconv.FormatFloat(v.Float(), 'f', -1, 32)
		break
	default:
		value = ""
		ok = false
		break
	}

	return value, ok
}

// InterfaceToInt 接口类型转整形
func InterfaceToInt(val interface{}, bitSize int) (int64, bool) {
	v := reflect.ValueOf(val)
	ok := true
	var value int64

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			value = 1
		} else {
			value = 0
		}

		break
	case reflect.String:
		if i, err := strconv.ParseInt(v.String(), 10, bitSize); err != nil {
			value = 0
		} else {
			value = i
		}
		break
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value = v.Int()
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value = int64(v.Uint())
		break
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value = int64(v.Float())
		break
	default:
		value = 0
		ok = false
		break
	}

	return value, ok
}

// InterfaceToUint 接口类型转无符号整形
func InterfaceToUint(val interface{}, bitSize int) (uint64, bool) {
	v := reflect.ValueOf(val)
	ok := true
	var value uint64

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			value = 1
		} else {
			value = 0
		}

		break
	case reflect.String:
		if i, err := strconv.ParseUint(v.String(), 10, bitSize); err != nil {
			value = 0
		} else {
			value = i
		}
		break
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value = uint64(v.Int())
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value = v.Uint()
		break
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value = uint64(v.Float())
		break
	default:
		value = 0
		ok = false
		break
	}

	return value, ok
}

// InterfaceToFloat 接口类型转浮点型
func InterfaceToFloat(val interface{}, bitSize int) (float64, bool) {
	v := reflect.ValueOf(val)
	ok := true
	var value float64

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			value = 1
		} else {
			value = 0
		}

		break
	case reflect.String:
		if i, err := strconv.ParseFloat(v.String(), bitSize); err != nil {
			value = 0
		} else {
			value = i
		}
		break
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value = float64(v.Int())
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value = float64(v.Uint())
		break
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value = v.Float()
		break
	default:
		value = 0
		ok = false
		break
	}

	return value, ok
}

// InterfaceToBool 接口类型转布尔类型
func InterfaceToBool(val interface{}) (bool, bool) {
	v := reflect.ValueOf(val)
	ok := true
	var value = true

	switch v.Kind() {
	case reflect.Bool:
		value = v.Bool()
		break
	case reflect.String:
		if v.String() == "" {
			value = false
		}
		break
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		if v.Int() <= 0 {
			value = false
		}
		break
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		if v.Uint() <= 0 {
			value = false
		}
		break
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		if v.Float() <= 0 {
			value = false
		}
		break
	default:
		value = false
		ok = false
		break
	}

	return value, ok
}

// InterfaceToStrings 接口类型转字符串数组
func InterfaceToStrings(val interface{}) ([]string, bool) {

	var res []string
	data, ok := val.([]interface{})
	if !ok {
		return nil, false
	}
	for _, v := range data {
		s, ok := InterfaceToString(v)
		if !ok {
			return nil, false
		}
		res = append(res, s)
	}
	return res, true
}
