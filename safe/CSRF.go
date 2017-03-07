package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login(w http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(time.Now().Unix(), 10))
		io.WriteString(h, "ganraomaxxxxxxxxx")
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)

	} else if request.Method == "POST" {
		request.ParseForm()
		token := request.Form.Get("token")
		fmt.Println(token)
		io.WriteString(w, token)
		fmt.Fprintf(w, token)
	}
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
