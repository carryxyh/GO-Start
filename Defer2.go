package main

import "fmt"

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func main() {
	fmt.Println(f())
	fmt.Println(test())
}

func f1() (result int) {
	result = 0 //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	func() {   //defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}

//从这个函数可看出，定义了返回类型，并声明了这个变量，则函数内部默认创建了这个变量
func test() (param int) {
	param = 1
	return param
}
