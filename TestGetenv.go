package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Getenv检索环境变量并返回值，如果变量是不存在的，这将是空的。
	HOME := os.Getenv("HOME")
	fmt.Println(HOME)
	//这里什么也不输出，应该获取的是系统中的 不是用户的
	fmt.Println(os.Getenv("JAVA_HOME"))
	//这里输出false  找不到
	fmt.Println(os.LookupEnv("JAVA_HOME"))
}
