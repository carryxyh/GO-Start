package main

import (
	"fmt"
)

func main() {
	v, s := test()
	fmt.Println(v)
	fmt.Println(s)
}

//只要v，s定义了  return会自动找到v，s的值，并且按照顺序返回。如果直接return则返回v和s的默认值
func test() (v int, s string) {
	//只注释v = 1 则打印 0 abc
	v = 1
	//下面这句注释，则只打印1。
	s = "abc"
	return
}
