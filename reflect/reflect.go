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

	fmt.Println("/*---------------分割线------------------*/")

	/*---------------分割线------------------*/

	testDynamicValue()

	fmt.Println("/*---------------分割线------------------*/")

	testMethodCall()
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
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

/*---------------分割线------------------*/

func testDynamicValue() {
	var num float64 = 3.1415

	//reflect.Value是通过reflect.ValueOf(X)获得的
	//只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值

	//如果传入的参数不是指针，而是变量，那么
	// 通过Elem获取原始值对应的对象则直接panic
	// 通过CanSet方法查询是否可以设置返回false
	pointer := reflect.ValueOf(&num)

	// reflect.Value.Elem() 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
	newValue := pointer.Elem()

	//3.1415
	fmt.Println("before set is ", newValue)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	fmt.Println("type of newValue ", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	newValue.SetFloat(6.28)

	// 这里会报错，严格控制类型
	// newValue.SetInt(123)
	fmt.Println("after set is ", newValue)
}

/*---------------分割线------------------*/

func (u User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name: ", name, ", age:", age, "and origal User.Name:", u.Name)
}

func (u User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func testMethodCall() {
	user := User{1, "ABC", 25}

	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value
	// 得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(user)

	//先看看带有参数的调用方法
	methodValue := getValue.MethodByName("ReflectCallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("111"), reflect.ValueOf(10)}
	methodValue.Call(args)

	methodValue = getValue.MethodByName("ReflectCallFuncNoArgs")
	args = make([]reflect.Value, 0)
	// 这里也是不行的。无参数必须是数组为0
	// args = make([]reflect.Value, 1)
	methodValue.Call(args)
}
