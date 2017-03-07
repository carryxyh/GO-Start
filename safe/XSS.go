package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	name := request.Form.Get("name")
	// s := template.HTMLEscape("login.gtpl")
	// fmt.Println(s)
	t, _ := template.ParseFiles("login.gtpl")
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", login)
	http.ListenAndServe(":7777", nil)
}
