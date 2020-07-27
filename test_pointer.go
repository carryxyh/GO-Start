package main

import (
	"fmt"
)

type P interface {
	PrintId()
}

type User struct {
	Id int64
}

func (t User) PrintId() {
	fmt.Println(t.Id)
}

func main() {
	u := &User{
		64,
	}

	xx(u)
}

func xx(in *User) {
	in.PrintId()
}
