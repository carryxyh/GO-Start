package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f1() (r int) {
	r = 1
	func(r int) { //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return r
}

func main() {
	fmt.Println(f())
}
