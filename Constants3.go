package main

import "fmt"
import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)

	UnKnown = 0
	Male    = 1
	FeMale  = 2
)

func main() {
	//这里如果修改了a的值，则a的值在println的时候会改变，但是b和c的值不会变，原理类似java的static
	a := "hello world"
	fmt.Println(a, b, c)
	fmt.Println(UnKnown)
}
