package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println(reflect.TypeOf(phone))
	fmt.Println(reflect.ValueOf(phone))

	v, ok := phone.(IPhone)
	//这里输出 &{}
	fmt.Println("v : ", v)
	//这里输出 true 如果换成phone.(IPHONE),输出false
	fmt.Println("ok : ", ok)

	var nokiaptr interface{}
	nokiaptr = &NokiaPhone{}

	realValue, o := nokiaptr.(NokiaPhone)

	fmt.Println(reflect.TypeOf(realValue))

	fmt.Println(realValue)
	fmt.Println(o)

	x := &realValue
	fmt.Println(&x)
	fmt.Println(&nokiaptr)
}
