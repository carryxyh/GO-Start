package main

import "fmt"

func main() {
	// var sz [10]int
	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var szbc = [...]int{1, 2}
	for i := 0; i < 2; i++ {
		fmt.Println(szbc[i])
	}
	//这里会报错，虽然是可变长的数组但是数组长度还是固定值
	fmt.Println(szbc[0])
}
