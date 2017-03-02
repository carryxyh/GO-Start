package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	//打开一个文件
	fl, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer fl.Close()

	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
