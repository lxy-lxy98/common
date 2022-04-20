package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// import (
// 	"fmt"
// 	"unsafe"
// )

// func main() {
// 	var i int = 0x12345678                  //定义数据
// 	const size int = int(unsafe.Sizeof(i))  //获取i的长度
// 	ps := (*[size]byte)(unsafe.Pointer(&i)) //size必须为const类型，不然会报错

// 	fmt.Printf("%T\n", ps) //*[8]byte类型
// 	fmt.Println(ps)
// 	//打印存储地址
// 	fmt.Println(&ps[0], ps[0])
// 	fmt.Println(&ps[1], ps[1])
// 	fmt.Println(&ps[2], ps[2])
// 	fmt.Println(&ps[3], ps[3])
// 	if ps[0] == 0x78 {
// 		//小端模式则ps[0]低地址存放的是低字节0x78，十进制为7*16+8=120,满足低地址存放低字节，存储为0x78563412不利用阅读，但方便计算机进行运算
// 		fmt.Println("系统为小端模式")
// 	} else {
// 		//大端模式则ps[0]低地址存放的是高字节0x12,十进制为18，满足低地址存放高字节，存储为0x12345678，方便阅读，但不方便计算机进行运算。
// 		fmt.Println("系统为大端模式")
// 	}
// }

type Slice []int

func (A Slice) Append(value int) {
	A1 := append(A, value)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&A))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n", sh.Data, sh.Len, sh.Cap)

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&A1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n", sh1.Data, sh1.Len, sh1.Cap)
}

func main() {
	mSlice := make(Slice, 10, 10)
	mSlice.Append(5)
	fmt.Println(mSlice)

}
