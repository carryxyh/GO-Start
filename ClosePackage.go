package main

import "fmt"

func main() {

	num1 := numAdd()

	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	num2 := numAdd()

	fmt.Println(num2())
	fmt.Println(num2())
}

func numAdd() func() int {
	var i int
	return func() int {
		i = i + 1
		return i
	}
}
