package main

import (
	"fmt"
	"reflect"
)

/**
在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型:
(value,type)
ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0
TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil
**/
func main() {
	var num float64 = 1.2345
	fmt.Println("type", reflect.TypeOf(num))
	fmt.Println("values", reflect.ValueOf(num))

	//print:
	//type float64
	//values 1.2345

	//reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
	//reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
	//也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	//强转
	//但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	//Golang 对类型要求非常严格，类型一定要完全符合
	//如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	//反射可以将“反射类型对象”再重新转换为“接口类型变量”
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	//这两行输出相同的值，convertPointer就相当于&num
	fmt.Println(convertPointer)
	fmt.Println(&num)

	//那么这里理所当然会输出num的值
	fmt.Println(*convertPointer)

	fmt.Println("------")

	fmt.Println(convertValue)
}
