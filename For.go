package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j <= 10; j++ {
		if j == 9 {
			continue
		}
		fmt.Println(j)
	}

	for {
		fmt.Println("I am break")
		break
	}

}
