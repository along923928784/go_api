package util

import (
	"fmt"
	"reflect"
	"strings"
)

func StructToMap1(obj interface{}) (classicData map[string]interface{}) {
	s := reflect.TypeOf(obj).Kind()
	fmt.Println(s)
	data := make(map[string]interface{})
	classicData = make(map[string]interface{})
	objV := reflect.ValueOf(obj).Elem()
	objT := objV.Type()
	for i := 0; i < objT.NumField(); i++ {
		name := strings.ToLower(objT.Field(i).Name)
		if name == "classicmodel" {
			classicData = ClassicStructToMap(objV.Field(i).Interface())
			continue
		}
		val := objV.Field(i).Interface()
		data[name] = val
	}
	for k, v := range data {
		classicData[k] = v
	}
	return
}

func ClassicStructToMap(obj interface{}) (data map[string]interface{}) {
	s := reflect.TypeOf(obj).Kind()
	fmt.Println(s)
	data = make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		tagVal := objT.Field(i).Tag.Get("json")
		if tagVal != "" {
			// name := objT.Field(i).Name
			val := objV.Field(i).Interface()
			data[tagVal] = val
		}

	}
	return
}

// 递归实现  性能待测试
func StructToMap(obj interface{}) (classicData map[string]interface{}) {
	var objV reflect.Value
	var objT reflect.Type
	data := make(map[string]interface{})
	classicData = make(map[string]interface{})

	objT = reflect.TypeOf(obj)
	objV = reflect.ValueOf(obj)
	if objT.Kind() == reflect.Ptr {
		objV = objV.Elem()
		objT = objV.Type()
	}

	for i := 0; i < objT.NumField(); i++ {
		name := strings.ToLower(objT.Field(i).Name)
		// if name == "basemodel" {
		// 	continue
		// }
		if name == "classicmodel" {
			classicData = StructToMap(objV.Field(i).Interface())
			continue
		}
		val := objV.Field(i).Interface()
		data[name] = val
	}
	for k, v := range data {
		classicData[k] = v
	}
	return
}
