package main

import "fmt"

func main() {
	// 旗标、宽度、精度、索引
	fmt.Printf("|%0+- #[1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 8, 4, 32, 64)

	// 浮点型精度
	fmt.Printf("|%f|%8.4f|%8.f|%.4f|%.f|\n", 3.2, 3.2, 3.2, 3.2, 3.2)
	fmt.Printf("|%.3f|%.3g|\n", 12.345678, 12.345678)
	fmt.Printf("|%.2f|\n", 12.345678+12.345678i)

	// 字符串精度
	s := "你好世界！"
	fmt.Printf("|%s|%8.2s|%8.s|%.2s|%.s|\n", s, s, s, s, s)
	fmt.Printf("|%x|%8.2x|%8.x|%.2x|%.x|\n", s, s, s, s, s)

	// 带引号字符串
	s1 := "Hello 世界!"       // CanBackquote
	s2 := "Hello\n世界!"      // !CanBackquote
	fmt.Printf("%q\n", s1)  // 双引号
	fmt.Printf("%#q\n", s1) // 反引号成功
	fmt.Printf("%#q\n", s2) // 反引号失败
	fmt.Printf("%+q\n", s2) // 仅包含 ASCII 字符

	// Unicode 码点
	fmt.Printf("%U, %#U\n", '好', '好')
	fmt.Printf("%U, %#U\n", '\n', '\n')

	// 接口类型将输出其内部包含的值
	var i interface{} = struct {
		name string
		age  int
	}{"AAA", 20}
	fmt.Printf("------------%v\n", i) // 只输出字段值
	fmt.Printf("%+v\n", i)            // 同时输出字段名
	fmt.Printf("%#v\n", i)            // Go 语法格式

	// 输出类型
	fmt.Printf("%T\n", i)
}
