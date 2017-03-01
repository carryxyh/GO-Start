package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	//如果这里email的 ‘e’ 是小写的，则下面输出不出来，但是不会报错
	Email string
}

func main() {
	t := template.New("fieldname example")
	// t, _ = t.Parse("hello {{.UserName}}")
	t, _ = t.Parse("hello {{.UserName}}! {{.Email}}")
	// p := Person{UserName: "ziyuan"}
	p := Person{UserName: "ziyuan", Email: "abc.com"}
	t.Execute(os.Stdout, p)
}
