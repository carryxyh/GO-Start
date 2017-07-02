package main

import (
	"fmt"
	"unsafe"
)

type Object struct {
	s string
}

func main() {
	obj := &Object{s: "1"}
	//unsafe.Pointer是特别定义的一种指针类型，它可以包含任意类型变量的地址,
	//我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型
	ptr := (*Object)(unsafe.Pointer(obj))
	fmt.Println(ptr)
	fmt.Println(ptr.s)
	// ref := int(ptr)
	// fmt.Println(ref)
}
