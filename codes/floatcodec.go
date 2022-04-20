package codes

import (
	"encoding/binary"
	"math"
	"reflect"
	"unsafe"
)

//ByteConvertFloat32  byte转换成float32
func ByteConvertFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes) //先将byte转换成uint32
	return math.Float32frombits(bits)         //将uint32转换成float32
}

//Float32ConvertByte  float32转换成byte
func Float32ConvertByte(float float32) []byte {
	bits := math.Float32bits(float) //float32转化成uint32
	bytes := make([]byte, 4)        //申请一个4字节的bytes
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

//ByteConvertFloat32Array
func ByteConvertFloat32Array(bf []byte, featureSize int) []float32 {
	features := make([]float32, featureSize) //初始化featureSIze大小的float32数组
	for i := 0; i < featureSize; i++ {
		off := i * 4
		features[i] = ByteConvertFloat32(bf[off : off+4]) //每四字节转换成一个float32数据
	}
	return features
}

//Float32ConvertByteArray
func Float32ConvertByteArray(values []float32) []byte {
	feature := []byte{}
	for _, value := range values {
		feature = append(feature, Float32ConvertByte(value)...)
	}
	return feature
}

//String 不会发生拷贝的[]byte转化string函数
func String(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

//Slice 不会发生拷贝的string转[]byte函数
func Slice(s string) (b []byte) {
	pstring := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}
