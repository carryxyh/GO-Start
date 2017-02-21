package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

type IPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (iphone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	v, ok := phone.(Phone)
	//这里输出 &{}
	fmt.Println("v : ", v)
	//这里输出 true 如果换成phone.(IPHONE),输出false
	fmt.Println("ok : ", ok)
}
