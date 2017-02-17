package main

import "fmt"
import "sync"
import "time"

//这是不会出问题的，go会检测运行环境中运行中的协程
func main() {
	var group sync.WaitGroup

	group.Add(1)

	go func() {
		time.Sleep(10 * time.Second)
		group.Done()
	}()

	group.Wait()
	fmt.Println("done!")
}
