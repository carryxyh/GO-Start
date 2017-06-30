package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	c := 3
	arr := make([]*int, 3)
	arr[0] = &a
	arr[1] = &b
	arr[2] = &c

	// fmt.Println(*arr)
	for i := range arr {
		fmt.Println(&arr[i])
	}
}
