package main

import "fmt"

func main() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//编译并不报错，运行时报错
	s := "txt"
	if s {
		fmt.Println("true")
	}

	//编译并不报错，运行时报错
	if nil {
		fmt.Println("nil true")
	}
}
