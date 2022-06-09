package Reflect

import "reflect"

// MapValueConversionSlice 将map的所有value放入slice中
// inMap: 类型map
// outSlice: 类型slice pointer
func MapValueConversionSlice(inMap interface{}, outSlice interface{}) {
	outSliceV := reflect.ValueOf(outSlice).Elem()
	inMapV := reflect.ValueOf(inMap)
	tempSlice := reflect.MakeSlice(outSliceV.Type(), 0, 0)
	for it := inMapV.MapRange(); it.Next(); {
		tempSlice = reflect.Append(tempSlice, it.Value())
	}
	outSliceV.Set(tempSlice)
}
