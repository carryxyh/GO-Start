package main

import "fmt"
import "sync"

var waitGroup sync.WaitGroup

func done(shut int) {
	fmt.Println(shut)
	waitGroup.Done()
}

func addJob() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
	}
}

func main() {
	//如果这里注释掉，会报错sync: negative WaitGroup counter
	//说明，要执行waitGroup.Done()，必须至少有一个goroutine在运行
	go addJob()

	//如果这里 i < 9 报错all goroutines are asleep - deadlock!
	for i := 0; i < 10; i++ {
		done(i)
	}

	waitGroup.Wait()
	fmt.Println("done!")
}
