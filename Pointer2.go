package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ptr *int
	var pptr **int

	ptr = &a

	pptr = &ptr

	fmt.Printf("指针变量 *ptr = %d\n", *ptr)          /* 这里应该输出10 */
	fmt.Printf("指向指针的指针变量 *pptr = %d\n", *pptr)   /* 这里输出ptr的地址 */
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) /* 这里输出10 */
}
