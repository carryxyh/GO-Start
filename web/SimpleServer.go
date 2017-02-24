package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/hello", sayHello) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9998", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)

	fmt.Println(r.Form)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println(r.Form["username"])
		fmt.Println(r.Form["userpwd"])
	}
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认不解析
	fmt.Println(r.Form) //这里发现 form是个键值对的形式
	fmt.Println("path : ", r.URL.Path)
	fmt.Println("scheme : ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("k : ", k)
		fmt.Println("val : ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello!")
}
