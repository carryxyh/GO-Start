package main

import "fmt"

// func main() {
// 	var a int = 10
// 	fmt.Println(&a)
// }

func main() {
	//在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
	//当一个指针被定义后没有分配到任何变量时，它的值为 nil。
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	ip = &a        /* 指针变量的值 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量的存储地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptr *int
	var ptrS *string

	fmt.Printf("ptr 的值为 : %x\n", ptr) /* 以下两个判断可以发现，指针的值为0，但是指针==nil */
	fmt.Println(ptr == nil)

	fmt.Printf("ptr 的值为 : %x\n", ptrS)
	fmt.Println(ptrS == nil)
}
