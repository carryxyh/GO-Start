package main

import (
	"fmt"
	"reflect"
)

func main() {

	user := UserA{"xyh", 1, func(food string) { fmt.Println("i eat ", food) }}

	DoFiledAndMethod(user)
}

// 通过接口来获取任意参数，然后一一揭晓
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	fmt.Println(getType.NumField())

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	fmt.Println(getType.NumMethod())

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

type UserA struct {
	Name string
	Id   int64
	Eat  func(food string)
}

// 如果使用 u *UserA 则打印不出来这个方法
func (u UserA) Phone() string {
	return u.Name
}
