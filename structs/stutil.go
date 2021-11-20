package structs

import (
	"encoding/json"
	"errors"
	"reflect"
)

// StructToMap struct转map
func StructToMap(obj interface{}) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	if obj == nil {
		err = errors.New("obj is nil")
		return
	}

	obj1 := reflect.TypeOf(obj)
	if obj1.Kind() != reflect.Struct {
		err = errors.New("type not Struct")
		return
	}
	obj2 := reflect.ValueOf(obj)

	for i := 0; i < obj1.NumField(); i++ {
		k := obj1.Field(i).Tag.Get("json")
		if k == "" {
			k = obj1.Field(i).Name
		}
		data[k] = obj2.Field(i).Interface()
	}
	return
}

// StructToJson 结构体转json
func StructToJson(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}

	obj := reflect.ValueOf(v)
	if obj.Kind() != reflect.Struct {
		return "", errors.New("value not Struct")
	}

	jsons, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsons), nil
}
