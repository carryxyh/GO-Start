package main

import (
	"fmt"
	"reflect"
)

type Sth interface {
	Say() int64
}

type User struct {
	id    int64
	Echo  func(sth int, name string) string
	useId func()
}

func (u User) Say() int64 {
	return u.id
}

func main() {
	s := &User{32, nil, nil}
	value := reflect.ValueOf(s)
	valueOfElem := value.Elem()
	fmt.Println(valueOfElem)

	uu := value.Interface().(*User)
	fmt.Println(uu)

	fields := valueOfElem.NumField()
	for i := 0; i < fields; i++ {
		field := valueOfElem.Field(i)
		if field.Kind() == reflect.Func && field.IsValid() && field.CanSet() {
			// is a public function
			funcType := field.Type()
			numOut := funcType.NumOut()
			outputs := make([]reflect.Value, numOut)

			field.Set(reflect.MakeFunc(field.Type(), func(input []reflect.Value) (output []reflect.Value) {
				for _, v := range input {
					inputValue := v.Interface()
					of := reflect.ValueOf(inputValue)
					if of.Kind() == reflect.String {
						x := "good man-" + inputValue.(string)
						outputs[0] = reflect.ValueOf(x)
					}
				}
				return outputs
			}))
		}
	}

	fmt.Println(s.Echo(1, "carryxyh"))
}
