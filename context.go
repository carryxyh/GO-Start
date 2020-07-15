package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	cancel()

	go handle(ctx, 500*time.Millisecond)

	time.After(10010 * time.Millisecond)

	select {
	case <-ctx.Done():
		xx := <-ctx.Done()
		fmt.Println(xx)
		fmt.Println("in main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
