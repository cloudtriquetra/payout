package utils

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	objMap := make(map[string]interface{})
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i).Interface()
		objMap[field.Name] = fieldValue
	}

	return objMap
}
