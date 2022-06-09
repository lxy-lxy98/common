package Reflect

import "reflect"

// MapKeyConversionSlice 将map的所有key放入slice中
// inMap: 类型map
// outSlice: 类型slice pointer
func MapKeyConversionSlice(inMap interface{}, outSlice interface{}) {
	outSliceV := reflect.ValueOf(outSlice).Elem()
	inMapV := reflect.ValueOf(inMap)
	outSliceV.Set(reflect.Append(outSliceV, inMapV.MapKeys()...))
}
