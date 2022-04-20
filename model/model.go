package model

import (
	"reflect"
	"time"
)

//PeriodicalExec 周期性传入方法 周期为period
func PeriodicalExec(method func(), period time.Duration) {
	go func() {
		for {
			method()
			time.Sleep(period)
		}
	}()
}

// CommonResultSlicePageLimit 通用的结果分页的方法
// offset: 截取切片的起始位置
// maxResult: 一次最多截取的结果数量
// inSlice: 要截取的切片，类型slice
// outSlice: 保存输出结果的切片，类型slice pointer
// return: 返回分页后下一页数据的起始位置，-1代表没有下一页
func CommonResultSlicePageLimit(offset int, maxResult int, inSlice interface{}, outSlice interface{}) int {
	inSliceV := reflect.ValueOf(inSlice)
	inSliceV.Type().Name()
	outSliceV := reflect.ValueOf(outSlice).Elem() ////Elem() 获取指针指向的值,
	if inSliceV.Len() < offset {
		offset = -1
	} else if inSliceV.Len() <= offset+maxResult {
		outSliceV.Set(inSliceV.Slice(offset, inSliceV.Len()))
		offset = -1
	} else {
		outSliceV.Set(inSliceV.Slice(offset, offset+maxResult))
		offset += maxResult
	}
	return offset
}
