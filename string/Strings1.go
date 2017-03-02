package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "od"))
	fmt.Println(strings.Count("abdab", "ab"))
	fmt.Println(strings.Fields("abcdef,"))
	//这里是不区分大小写的  都为true
	fmt.Println(strings.EqualFold("abc", "ABc"))
	//比较大小,应该世ASCII码比较
	fmt.Println(strings.Compare("abcd", "c"))

	//字符串链接
	s := []string{"a", "b", "ziyuan"}
	//这里join的参数是个slice
	fmt.Println(strings.Join(s, "0.0"))

	//查找字符串
	fmt.Println(strings.Index("ziyuan", "z"))
	//找不到这里返回-1，跟java用法一样
	fmt.Println(strings.Index("ziyuan", "123"))

	//重复字符串
	fmt.Println(strings.Repeat("na", 2))

	//在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("ziyuan", "z", "X", 1))

	//把s字符串按照sep分割，返回slice
	fmt.Println(strings.Split("zi,yuan", ",")[1])

	//在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Println(strings.Trim(" !!! ziyuan ?!", "!"))

	//去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Println(strings.Fields("zi yuan shi wu liang "))
}
