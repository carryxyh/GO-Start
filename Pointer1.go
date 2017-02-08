package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{1, 2, 3}
	var ptr [MAX]*int

	for i := 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i := 0; i < MAX; i++ {
		fmt.Println(*ptr[i])
	}
}
