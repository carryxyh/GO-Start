package main

import (
	"fmt"
	"os"
)

//这里写入是覆盖的写，不是追加，追加要用 func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
func main() {
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test \n")
		fout.Write([]byte("test!!!"))
	}
}
