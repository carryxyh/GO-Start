package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值,但是main函数中的a和b的值没有发生变化,说明是值传递 */
	swap(a, b)

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)
}

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}
