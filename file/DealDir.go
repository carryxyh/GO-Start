package main

import (
	"os"
)

func main() {
	os.Mkdir("./test", 0777)
	os.MkdirAll("./test/test1", 0777)
	err := os.Remove("./test")
	if err != nil {
		//ignore
	}
	// os.RemoveAll("./test")
}
