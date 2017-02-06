package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

// func main() {
// 	fmt.Println(a, b, c)
// }

// func main() {
// 	const (
// 		a = iota
// 		b
// 		c
// 		d = "ha" //独立值，iota += 1
// 		e        //"ha"   iota += 1
// 		f = 100  //iota +=1
// 		g        //100  iota +=1
// 		h = iota
// 		i
// 	)
// 	fmt.Println(a, b, c)
// 	fmt.Println(a, b, c, d, e, f, g, h, i)
// 	fmt.Println(1 << 0)
// 	fmt.Println(3 << 1)
// 	fmt.Println(4 << 1)
// }

func main() {
	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i, j, k, l)

	//这里的例子说明，如果是const中进行了iota的运算，则运算的最后一次过程会继续往下传递
	const (
		c = 1 + iota //iota = 0
		n = 2 + iota //iota = 1
		x            //iota = 2，这里x = 2 + 2
		y            //iota = 3，这里y = 2 + 3
		z = 3 + iota //iota = 4，z = 3 + 4 = 7
		p            //iota = 5 ，p = 3 + 5 = 8
	)
	fmt.Println(c, n, x, y, z, p)
}
