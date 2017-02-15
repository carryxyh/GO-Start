package main

import (
	"fmt"
	// "time"
)

func main() {
	channel := make(chan string)

	go func() {
		// fmt.Println("sleep 1")
		channel <- "wo ri ni ma"
		// fmt.Println("sleep 2")
	}()

	//这里输出 wo ri ni ma，但是注意！如果这里从channel中读出了数据，后面会报错，因为channel中已经没有数据了
	fmt.Println(<-channel)

	//这里仍然报错，why？！
	msg1 := <-channel
	fmt.Println(msg1)
}
