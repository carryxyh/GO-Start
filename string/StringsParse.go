package main

import (
	"fmt"
	"strconv"
)

func main() {
	//make([]type, len, cap)
	str := make([]byte, 0, 100)

	//Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中,最后一个参数应该是 “进制 10就是十进制”
	str = strconv.AppendInt(str, 4567, 10)
	fmt.Println(str)

	str = strconv.AppendBool(str, false)
	fmt.Println(str)

	str = strconv.AppendQuote(str, "abcdefg")
	fmt.Println(str)

	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(str)

	//这里输出4567false"abcdefg"'单'，可以看到 上面是把一些值改成了ascii码存下来，然后输出时要用string（str）
	fmt.Println(string(str))

	/*--------------------------------------------*/

	//Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	fmt.Println(a)

	b := strconv.FormatFloat(123.12, 'g', 12, 64)
	fmt.Println(b)

	//第二个参数同样是进制
	d := strconv.FormatUint(12345, 3)
	fmt.Println(d)

	//Itoa是FormatInt（int64（i），10）的缩写。
	e := strconv.Itoa(1025)
	fmt.Println(e)

	/*--------------------------------------------*/

	//Parse 系列函数把字符串转换为其他类型
	x, err := strconv.ParseBool("false")
	checkError(err)
	fmt.Println(x)

	y, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	fmt.Println(y)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
