package main

import (
	"fmt"
)

func main() {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
}

func a() int {
	var i int
	fmt.Println("外面a addr:", &i) // 打印结果为 a defer1: 1
	defer func() {
		fmt.Println(i)
	}()
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
		fmt.Println("a addr:", &i)  // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	// 这里return的i并不是真正的i
	// 我们可以认为是
	// 1. x=i
	// 2. return x
	// 只是把i的值返回了，所以defer中无法直接修改x的值
	return i
}
