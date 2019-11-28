package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Map2StructOption 选项结构
type Map2StructOption struct {
	KeyNameTagName      string
	DefaultTagName      string
	TimeFormatTagName   string
	TimeLocationTagName string
}

// Map2Struct map转struct
func Map2Struct(m map[string]interface{}, stc interface{}) (err error) {
	return Map2StructWithOption(m, reflect.ValueOf(stc), &Map2StructOption{
		KeyNameTagName:      "json",
		DefaultTagName:      "default",
		TimeFormatTagName:   "time_format",
		TimeLocationTagName: "time_location",
	})
}

// Map2StructWithOption map按选项转struct
func Map2StructWithOption(m map[string]interface{}, stc interface{}, option *Map2StructOption) (err error) {
	if option.KeyNameTagName == "" {
		option.KeyNameTagName = "json"
	}
	if option.DefaultTagName == "" {
		option.DefaultTagName = "default"
	}
	if option.TimeFormatTagName == "" {
		option.TimeFormatTagName = "time_format"
	}
	if option.TimeLocationTagName == "" {
		option.TimeLocationTagName = "time_location"
	}
	return map2Struct(m, reflect.ValueOf(stc), option)
}

// map2Struct map转struct
func map2Struct(m map[string]interface{}, val reflect.Value, option *Map2StructOption) (err error) {
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("paramter error")
	}

	vType := val.Type()

	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		mapKey := field.Tag.Get(option.KeyNameTagName)
		err := convertAssign(val.Field(i), field, m[mapKey], option)
		if err != nil {
			return err
		}
	}
	return nil
}

// convertAssign
// destValue 需要赋值的字段Value
// destFiel 需要赋值的字段StructField
// src 需要赋值的数据
func convertAssign(destValue reflect.Value, destField reflect.StructField, src interface{}, option *Map2StructOption) error {

	// 获取默认值
	defaulValue := destField.Tag.Get(option.DefaultTagName)

	// 判断需要赋值的类型
	k := destValue.Kind()
	if isSimpleKind(k) {
		return convertAssignSimpleKind(src, destValue, defaulValue)
	} else if isListKind(k) {
		return convertAssignListKind(src, destValue, destField, option)
	} else if isStructKind(k) {
		return convertAssignStructKind(src, destValue, destField, option)
	}

	// 其他类型
	switch destValue.Kind() {
	case reflect.Ptr:
		//return convertAssignPtrKind(src, destValue)
		return fmt.Errorf("不支持struct中存在指针类型字段：%v", destField.Name)
	case reflect.Map:
		return fmt.Errorf("不支持struct中存在map类型字段：%v", destField.Name)
	}

	return fmt.Errorf("不支持的类型字段：%v", destField.Name)
}

// isSimpleKind 判断是否是简单类型，数字、字符、布尔等
func isSimpleKind(k reflect.Kind) bool {
	if k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64 ||
		k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64 ||
		k == reflect.Float32 ||
		k == reflect.Float64 ||
		k == reflect.String ||
		k == reflect.Bool {
		return true
	}
	return false
}

// isListKind 判断是否是数组、切片
func isListKind(k reflect.Kind) bool {
	if k == reflect.Array || k == reflect.Slice {
		return true
	}
	return false
}

// isStructKind 判断是否是结构
func isStructKind(k reflect.Kind) bool {
	if k == reflect.Struct {
		return true
	}
	return false
}

// convertAssignSimpleKind 简单类型值赋值
func convertAssignSimpleKind(s interface{}, v reflect.Value, def string) error {
	switch v.Kind() {
	case reflect.Int:
		setInt(s, v, 0, def)
	case reflect.Int8:
		setInt(s, v, 8, def)
	case reflect.Int16:
		setInt(s, v, 16, def)
	case reflect.Int32:
		setInt(s, v, 32, def)
	case reflect.Int64:
		setInt(s, v, 64, def)
	case reflect.Uint:
		setUint(s, v, 0, def)
	case reflect.Uint8:
		setUint(s, v, 8, def)
	case reflect.Uint16:
		setUint(s, v, 16, def)
	case reflect.Uint32:
		setUint(s, v, 32, def)
	case reflect.Uint64:
		setUint(s, v, 64, def)
	case reflect.Float32:
		setFloat(s, v, 32, def)
	case reflect.Float64:
		setFloat(s, v, 64, def)
	case reflect.String:
		setString(s, v, def)
	case reflect.Bool:
		setBool(s, v, def)
	}
	return nil
}

// convertAssignListKind 数组、切片类型值赋值
func convertAssignListKind(s interface{}, v reflect.Value, f reflect.StructField, option *Map2StructOption) error {
	// 判断数组、切片内部数据类型
	elemKind := v.Type().Elem().Kind()

	if isSimpleKind(elemKind) {
		setListSimpleKind(s, v)
	} else if isStructKind(elemKind) {
		setListStructKind(s, v, f, option)
	} else if isListKind(elemKind) {
		// TODO:
		return fmt.Errorf("暂不支持数组中嵌套数组")
	}
	return nil
}

// convertAssignStructKind 结构类型值赋值
func convertAssignStructKind(s interface{}, v reflect.Value, f reflect.StructField, option *Map2StructOption) error {
	switch v.Interface().(type) {
	case time.Time:
		setTime(s, v, f, option)
		return nil
	}
	map2Struct(s.(map[string]interface{}), v, option)
	return nil
}

// setInt 设置int
func setInt(s interface{}, v reflect.Value, bitSize int, def string) {
	intVal, _ := InterfaceToInt(s, bitSize)

	if intVal == 0 {
		intVal, _ = strconv.ParseInt(def, 10, bitSize)
	}
	v.SetInt(intVal)
}

// setUint 设置uint
func setUint(s interface{}, v reflect.Value, bitSize int, def string) {
	uintVal, _ := InterfaceToUint(s, bitSize)

	if uintVal == 0 {
		uintVal, _ = strconv.ParseUint(def, 10, bitSize)
	}
	v.SetUint(uintVal)
}

// setFloat 设置float
func setFloat(s interface{}, v reflect.Value, bitSize int, def string) {
	floatVal, _ := InterfaceToFloat(s, bitSize)

	if floatVal == 0 {
		floatVal, _ = strconv.ParseFloat(def, bitSize)
	}

	v.SetFloat(floatVal)
}

// setString 设置string
func setString(s interface{}, v reflect.Value, def string) {
	var stringVal string
	stringVal, _ = InterfaceToString(s)

	if stringVal == "" {
		stringVal = def
	}
	v.SetString(stringVal)
}

// setBool 设置bool
func setBool(s interface{}, v reflect.Value, def string) {
	boolVal, _ := InterfaceToBool(s)
	v.SetBool(boolVal)
}

// setListSimpleKind 数组、切片中项的类型为简单类型的赋值
func setListSimpleKind(s interface{}, v reflect.Value) {
	switch v.Type().Elem().Kind() {
	case reflect.Int:
		data := make([]int, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			intVal, _ := InterfaceToInt(sValue[i], 0)
			newData = append(newData, reflect.ValueOf(int(intVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Int8:
		data := make([]int8, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			intVal, _ := InterfaceToInt(sValue[i], 8)
			newData = append(newData, reflect.ValueOf(int8(intVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Int16:
		data := make([]int16, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			intVal, ok := InterfaceToInt(sValue[i], 16)
			if !ok {
				intVal = 0
			}
			newData = append(newData, reflect.ValueOf(int16(intVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Int32:
		data := make([]int32, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			intVal, _ := InterfaceToInt(sValue[i], 32)
			newData = append(newData, reflect.ValueOf(int32(intVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Int64:
		data := make([]int64, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			intVal, _ := InterfaceToInt(sValue[i], 64)
			newData = append(newData, reflect.ValueOf(intVal))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Uint:
		data := make([]uint, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			uintVal, _ := InterfaceToUint(sValue[i], 0)
			newData = append(newData, reflect.ValueOf(uint(uintVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Uint8:
		data := make([]uint8, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			uintVal, _ := InterfaceToUint(sValue[i], 8)
			newData = append(newData, reflect.ValueOf(uint8(uintVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Uint16:
		data := make([]uint16, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			uintVal, _ := InterfaceToUint(sValue[i], 16)
			newData = append(newData, reflect.ValueOf(uint16(uintVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Uint32:
		data := make([]uint32, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			uintVal, _ := InterfaceToUint(sValue[i], 32)
			newData = append(newData, reflect.ValueOf(uint32(uintVal)))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Uint64:
		data := make([]uint64, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			uintVal, _ := InterfaceToUint(sValue[i], 64)
			newData = append(newData, reflect.ValueOf(uintVal))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Float32:
		data := make([]float32, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			floatVal, _ := InterfaceToFloat(sValue[i], 32)
			newData = append(newData, reflect.ValueOf(floatVal))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Float64:
		data := make([]float64, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			floatVal, _ := InterfaceToFloat(sValue[i], 64)
			newData = append(newData, reflect.ValueOf(floatVal))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.String:
		data := make([]string, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			str, _ := InterfaceToString(sValue[i])
			newData = append(newData, reflect.ValueOf(str))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	case reflect.Bool:
		data := make([]bool, 0)
		dataValue := reflect.ValueOf(data)
		newData := make([]reflect.Value, 0)
		sValue := s.([]interface{})
		for i := 0; i < len(sValue); i++ {
			boolVal, _ := InterfaceToBool(sValue[i])
			newData = append(newData, reflect.ValueOf(boolVal))
		}

		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	}
}

// setListSimpleKind 数组、切片中项的类型为结构类型的赋值
func setListStructKind(s interface{}, v reflect.Value, f reflect.StructField, option *Map2StructOption) {
	cval := reflect.New(v.Type().Elem()).Elem()
	switch cval.Interface().(type) {
	case time.Time:
		dataValue := reflect.ValueOf(v.Interface())
		newData := make([]reflect.Value, 0)
		sData := s.([]interface{})
		for _, data := range sData {
			setTime(data, cval, f, option)
			newData = append(newData, reflect.ValueOf(cval.Interface()))
		}
		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	default:
		dataValue := reflect.ValueOf(v.Interface())
		newData := make([]reflect.Value, 0)
		sData := s.([]interface{})
		for _, data := range sData {
			map2Struct(data.(map[string]interface{}), cval, option)
			newData = append(newData, reflect.ValueOf(cval.Interface()))
		}
		val := reflect.Append(dataValue, newData...)
		v.Set(val)
	}
}

// setTime 设置time
func setTime(s interface{}, v reflect.Value, field reflect.StructField, option *Map2StructOption) {
	timeFormat := field.Tag.Get(option.TimeFormatTagName)
	if timeFormat == "" {
		timeFormat = time.RFC3339
	}

	timeLocation := field.Tag.Get(option.TimeLocationTagName)
	local := time.Local
	if timeLocation != "" {
		loc, err := time.LoadLocation(timeLocation)
		if err != nil {
			return
		}
		local = loc
	}
	var timeVal time.Time
	var err error

	switch s.(type) {
	case string:
		timeVal, err = time.ParseInLocation(timeFormat, s.(string), local)
	default:
		timeVal, err = time.ParseInLocation(timeFormat, s.(string), local)
	}

	if err == nil {
		v.Set(reflect.ValueOf(timeVal))
	}
}
