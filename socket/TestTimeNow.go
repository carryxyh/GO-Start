package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	daytime := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(daytime)
}
