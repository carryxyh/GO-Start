package main

import (
	"context"
	"fmt"
	"time"
)

// 构建一个父context，用来取消子context
func main() {

	parent, cancelFunc := context.WithCancel(context.Background())
	// 由于parent在handle中500ms之后就被cancel了，所以child也会被cancel，实际上child done在500ms之后就会被触发，不会等10000ms
	child, cancel := context.WithTimeout(parent, 10000*time.Millisecond)
	defer cancel()

	go handle(cancelFunc, 500*time.Millisecond)

	select {
	case <-child.Done():
		if child.Err() != nil {
			fmt.Println("has err?", child.Err())
		}
		fmt.Println("main over")
	}
}

func handle(cancelFunc func(), duration time.Duration) {
	<-time.After(duration)
	fmt.Println("cancel child")
	cancelFunc()
}
