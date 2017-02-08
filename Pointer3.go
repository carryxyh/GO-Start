package main

import "fmt"

func main() {
	var arr = [2]int{1, 2}
	var ptr *int

	ptr = &arr[0]
	fmt.Println(*ptr)

	var ptrj *int
	//这里会报错，说明go语言的指针没法像C一样进行加减操作
	ptrj = ptr + 1
	fmt.Println(*ptrj)
}
