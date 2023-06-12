package gstru

import (
	"encoding/json"
	"reflect"
	"strings"
)

/*
   @File: stru.go
   @Author: khaosles
   @Time: 2023/3/5 12:26
   @Desc:
*/

// Assignment 把结构体1的相同字段赋值给结构体2
func Assignment[T1, T2 any](elem1 T1, elem2 *T2) {
	// 获取结构体类型和值
	t1 := reflect.TypeOf(elem1)
	v1 := reflect.ValueOf(elem1)
	t2 := reflect.TypeOf(*elem2)
	v2 := reflect.ValueOf(elem2).Elem()

	// 遍历1结构体的字段
	for i := 0; i < t1.NumField(); i++ {
		// 获取字段名称和值
		fieldName := t1.Field(i).Name
		fieldValue := v1.Field(i).Interface()
		// 判断2结构体是否有同名字段
		if _, ok := t2.FieldByName(fieldName); ok {
			// 根据字段名称设置2结构体相应字段的值
			v2.FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}
}

// StructToMap struct to map
func StructToMap(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	// 如果传入的不是结构体指针，则直接返回空 map
	if objType.Kind() != reflect.Ptr || objType.Elem().Kind() != reflect.Struct {
		return map[string]string{}
	}

	data := make(map[string]string)
	for i := 0; i < objValue.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)
		value := objValue.Elem().Field(i)

		// 如果字段是空值，则跳过
		if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			jsonTags := strings.Split(jsonTag, ",")
			name := strings.TrimSpace(jsonTags[0])
			data[name] = toString(value.Interface())
		} else {
			data[field.Name] = toString(value.Interface())
		}

		// 如果字段是结构体类型，则递归调用 StructToMap 进行转换
		if field.Type.Kind() == reflect.Struct {
			nestedData := StructToMap(value.Interface())
			for k, v := range nestedData {
				data[k] = v
			}
		}
	}

	return data
}

func StructToMapInterface(obj interface{}) map[string]any {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	// 如果传入的不是结构体指针，则直接返回空 map
	if objType.Kind() != reflect.Ptr || objType.Elem().Kind() != reflect.Struct {
		return map[string]any{}
	}
	data := make(map[string]any)
	for i := 0; i < objValue.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)
		value := objValue.Elem().Field(i)
		data[field.Name] = value.Interface()
	}

	return data
}

func toString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func SetField(obj interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(obj).Elem() // 获取结构体的反射值
	f := v.FieldByName(fieldName)    // 获取属性的反射值
	if f.IsValid() && f.CanSet() {   // 检查属性是否存在且可设置
		val := reflect.ValueOf(value) // 获取要赋的值的反射值
		if f.Type() == val.Type() {   // 检查值的类型是否与属性类型匹配
			f.Set(val) // 给属性赋值
		}
	}
}
