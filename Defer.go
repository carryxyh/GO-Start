package main

import "fmt"

func main() {

	// defer fmt.Println(1)

	// return

	// defer fmt.Println(2)

	// a()
	// fmt.Println(b(1, "2"))

	var test int
	test = 1
	defer fmt.Println("第一个defer")
	if test == 2 {
		return
	}
	defer fmt.Println("第二个defer")
}

func a() {
	if true {
		return
	}
	defer fmt.Println(1)
}

func b(x int, y string) (int, string) {
	//这样是不行的
	// return
	if x == 1 {
		//这里也是会报错的  要写成 return x,y
		return
	}
	return x, y
}
