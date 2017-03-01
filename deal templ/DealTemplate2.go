package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "zzz"}
	f2 := Friend{Fname: "ziy"}

	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.UserName}}!
                {{range .Emails}}
                    an email {{.}}
                {{end}}
                {{with .Friends}}
                {{range .}}
                    my friend name is {{.Fname}}
                {{end}}
                {{end}}
                `)

	p := Person{UserName: "ziyuan",
		Emails:  []string{"z1", "z2"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
