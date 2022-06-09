package Reflect

import "reflect"

func CheckType(v interface{}) reflect.Kind {
	t := reflect.TypeOf(v)
	return t.Kind()
}
