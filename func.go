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

	// 这里这个add方法相当于handler方法的一个属性，调用时需要传一个handler方法实例（因为这里的h handler 实际上是个声明而不是定义）
	fmt.Println(handler.add(func(name string) int {
		return 10
	}, "1"))

	fmt.Println(filter(isTen, 10))
	fmt.Println(filter(isNotTen, 10))
}

type condition func(int) bool

// 这里的isTen可以作为condition的实例来使用 任何一个入参是int 返回是bool的函数都可以
func isTen(param int) bool {
	return param == 10
}

func isNotTen(param int) bool {
	return param != 10
}

func filter(c condition, param int) bool {
	return c(param)
}
