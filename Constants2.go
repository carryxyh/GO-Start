package main

import "fmt"

func main() {
	const LENGTH = 10
	const WIDTH = 5

	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH

	fmt.Printf("面积 %d", area)
	println()
	println(a, b, c)
}
