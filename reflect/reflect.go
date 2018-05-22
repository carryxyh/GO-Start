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

	//这里同样输出1.2345
	fmt.Println("print-----", value.Interface())

	//这两行输出相同的值，convertPointer就相当于&num
	fmt.Println(convertPointer)
	fmt.Println(&num)

	//那么这里理所当然会输出num的值
	fmt.Println(*convertPointer)

	fmt.Println("/*---------------分割线------------------*/")

	fmt.Println(convertValue)

	/*---------------分割线------------------*/

	testField()
}

/*---------------分割线------------------*/

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("ReflectCallFunc")
}

func testField() {
	user := User{1, "ABC", 25}

	getType := reflect.TypeOf(user)

	fmt.Println("getType is", getType.Name())

	getValue := reflect.ValueOf(user)

	fmt.Println("getValue is ", getValue)

	getPointer := reflect.TypeOf(&user)

	//*main.User
	fmt.Println("getPointer is ", getPointer)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {

		//注意这两条，第一个要获取的是字段类型的相关信息
		//第二个是获取字段值相关的信息
		field := getType.Field(i)
		value := getValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
}
