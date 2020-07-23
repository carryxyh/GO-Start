package main

import (
	"fmt"
	"reflect"
)

func main() {
	inter := EchoService{}
	valueOf := reflect.ValueOf(inter)
	fmt.Println(valueOf)

}

type Service interface {
	Reference() string
}

type User struct {
	name string
	id   int64
}

type EchoService struct {
	Echo           func(value string) string
	EchoComplicate func(user User) User
}

func (echoService *EchoService) Reference() string {
	return "echoService"
}
