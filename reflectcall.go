package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int { return a + b }

func main() {

	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}

	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Interface()) // #=> 1
}
