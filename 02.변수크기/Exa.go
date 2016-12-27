// Exa
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i1 int8
	fmt.Println("int8 크기:", unsafe.Sizeof(i1), "최소값:", -0x80, "최대값:", 0x7F)
	var i2 int16
	fmt.Println("int16 크기:", unsafe.Sizeof(i2), "최소값:", -0x8000, "초대값:", 0x7FFF)
	var i3 int32
	fmt.Println("int32 크기:", unsafe.Sizeof(i3), "최소값:", -0x80000000, "최대값:", 0x7FFFFFFF)
	var i4 int64
	fmt.Println("int64 크기:", unsafe.Sizeof(i4), "최소값:", -0x8000000000000000, "최대값:", 0x7FFFFFFFFFFFFFFF)
}
