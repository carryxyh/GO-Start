package main

import (
	"fmt"
)

/*
	go语言中，interface可以转化为任何一个类型。
	var s interface{} = "abc"  相当于java中的  Object s = "abc"
*/
func main() {
	var s interface{} = "abc"
	j, ok := s.(int)
	fmt.Println(j)
	fmt.Println(ok)
}
