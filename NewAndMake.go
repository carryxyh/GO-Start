package main

import "fmt"

func main() {
	var p *[]int = new([]int)
	//这里是false
	fmt.Println(p == nil)

	*p = make([]int, 1, 2)
	//这里输出 &[0]
	fmt.Println(p)

	v := make([]int, 1)
	//这里输出 [0]
	fmt.Println(v)
	//这里也输出 &[0]
	fmt.Println(&v)
}
