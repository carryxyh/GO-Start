package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		fmt.Println("func value")
		r = r + 5
	}(r)

	//相比于Defer4的实例，这里加了一个指针的，输出是
	//func ptr
	//func value
	//说明defer 越后面就越先执行
	defer func(ptr *int) {
		fmt.Println("func ptr")
		*ptr = *ptr + 5
	}(&r)
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
