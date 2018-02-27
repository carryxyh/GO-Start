package main

import (
	"fmt"
)

//定义函数类型
type handler func(name string) int

//实现函数类型方法
func (h handler) add(name string) int {
	return h(name) + 10
}

func decrease(h handler, param int) int {
	return h("1") - param
}

func main() {
	fmt.Println(decrease(func(name string) int {
		return 10
	}, 5))

	fmt.Println(handler.add(func(name string) int {
		return 10
	}, "1"))
}
