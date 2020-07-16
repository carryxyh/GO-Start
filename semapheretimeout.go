package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	sema := semaphore.NewWeighted(10)
	er1 := sema.Acquire(ctx, 5)
	if er1 != nil {
		fmt.Println("err1")
	}
	er2 := sema.Acquire(ctx, 6)
	if er2 != nil {
		// 打印err2，因为获取超时了，ctx会在5s之后超时，此时会正常终结acquire
		fmt.Println("err2")
	}

	fmt.Println("main over")
}
