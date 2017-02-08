package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 6)
	printSlice(numbers)
}

func printSlice(sli []int) {
	fmt.Printf("len = %d , cap = %d , slice = %v \n", len(sli), cap(sli), sli)
}
