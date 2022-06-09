package Reflect

import (
	"errors"
	"reflect"
)

func FillBySetting(st interface{}, settings map[string]interface{}) error {
	//首先需要判断st的类型，因为需要修改其中的值，所以是指针类型
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer")
	}
	//st需要是结构体类型的指针 st类型需要是Array, Chan, Map, Ptr或Slice之一，否则调用Kind()会panic
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}
	if settings == nil {
		return errors.New("settings is nil")
	}
	var (
		field reflect.StructField //结构中的单个字段名
		ok    bool
	)
	for k, v := range settings {
		//先判断key是否存在，再判断key的类型是否与结构体重key类型一致
		if field, ok = reflect.ValueOf(st).Elem().Type().FieldByName(k); !ok { //如果传入结构体指针中没有key，continue
			continue
		}
		//填充
		if field.Type == reflect.TypeOf(v) { //如果结构体其中一个key类型与值类型一致
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}
