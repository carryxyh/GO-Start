package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var x struct {
		a bool
		b int16
		c []int
	}

	/**
	  unsafe.Offsetof 函数的参数必须是一个字段 x.f, 然后返回 f 字段相对于 x 起始地址的偏移量, 包括可能的空洞.
	*/

	/**
	  uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	  指针的运算
	*/
	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b) // "42"
}

//上面的写法非常繁琐，功能：
//将变量x的地址加上b字段地址偏移量转化为*int16类型指针，然后通过该指针更新x.b：

//下面段代码是错误的：
func main() {
	// NOTE: subtly incorrect!
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
}

//产生错误的原因很微妙。有时候垃圾回收器会移动一些变量以降低内存碎片等问题。
//这类垃圾回收器被称为移动GC。当一个变量被移动，所有的保存改变量旧地址的指针必须同时被更新为变量移动后的新地址。
//从垃圾收集器的视角来看，一个unsafe.Pointer是一个指向变量的指针，因此当变量被移动是对应的指针也必须被更新；
//但是uintptr类型的临时变量只是一个普通的数字，所以其值不应该被改变。
//上面错误的代码因为引入一个非指针的临时变量tmp，导致垃圾收集器无法正确识别这个是一个指向变量x的指针。
//当第二个语句执行时，变量x可能已经被转移，这时候临时变量tmp也就不再是现在的&x.b地址。
//第三个向之前无效地址空间的赋值语句将彻底摧毁整个程序！
