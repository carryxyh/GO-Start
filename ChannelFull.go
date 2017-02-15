package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "123"
	messages <- "abc"

	//使用select，如果能成功，放入，否则会输出messages is full
	select {

	case messages <- "456":
		fmt.Println("ok")
	default:
		fmt.Println("messages is full")
	}
	// messages <- "堵住！"
	fmt.Println("...")
	fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
}

// func main() {
// 	channel := make(chan string, 2)

// 	channel <- "h1"

// 	channel <- "w2"

// 	select {

// 	case channel <- "c3":
// 		fmt.Println("ok")
// 	default:
// 		fmt.Println("channel is full !")
// 	}

// 	fmt.Println("...")
// 	// msg1 := <-channel
// 	// fmt.Println(msg1)
// }
