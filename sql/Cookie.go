package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//设置cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, cookie)

	//读取cookie方法1
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)

	//读取cookie方法2
	for _, cookie := range r.Cookies() {
		fmt.Fprint(w, cookie.Name)
	}
}
