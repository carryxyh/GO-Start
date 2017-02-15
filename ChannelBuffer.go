package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "123"
	messages <- "abc"
	messages <- "aaaaa"
	//如果有这句，会报错 all goroutines are asleep - deadlock!
	messages <- "堵住！"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
}
