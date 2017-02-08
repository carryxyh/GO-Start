package main

import "fmt"

func main() {
	fmt.Println(re(1, "2 -> string"))
}

func re(x int, y string) (int, string) {
	return x, y
}
