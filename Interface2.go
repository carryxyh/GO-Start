package main

import "fmt"

type geometry interface {
	area() float64
	zc() float64
}

type fangxing struct {
	width  float64
	height float64
}

type circle struct {
	r float64
}

func (fx fangxing) area() float64 {
	return fx.height * fx.width
}

func (fx fangxing) zc() float64 {
	return 2*fx.height + 2*fx.width
}

func (yx circle) zc() float64 {
	return 2 * 3.14 * yx.r
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.zc())
}

func (yx circle) area() float64 {
	return 3.14 * yx.r * yx.r
}

func main() {
	fx := fangxing{width: 3, height: 4}
	yx := circle{r: 2}

	//这里是会报错的 由于声明了一个geometry的变量，其中没有width和height
	// var fx geometry
	// fx = new(fangxing)
	// fx.width = 3
	// fx.height = 4

	//这里也报错了，fx不是一个 *fangxing的type，根据interface.go中的写法来看。。问题出在 var fx fangxing
	// var fx fangxing
	// fx = new(fangxing)
	// fx.width = 3
	// fx.height = 4

	measure(fx)
	measure(yx)
}
