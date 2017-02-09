package main

import "fmt"

//return 语句可以拆分成这么几个步骤
//返回值 = xxx
//调用defer函数
//空的return
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func main() {
	fmt.Println(f())
}

//所以f函数会变成这样
func f1() (r int) {
	t := 5
	r = t    //赋值指令
	func() { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
		t = t + 5
	}
	return //空的return指令
}
