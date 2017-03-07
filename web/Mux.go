package main

import (
	"fmt"
	"net/http"
)

//自定义的路由器
type MyMux struct {
}

func (m *MyMux) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}

	http.NotFound(w, r)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
