package main

import (
	"fmt"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delegating...")
	//会先跑到这里，再跑到下面的viewRecord
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func viewRecord(w http.ResponseWriter, r *http.Request) error {
	// c := appengine.NewContext(r)
	// key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	// record := new(Record)
	// if err := datastore.Get(c, key, record); err != nil {
	// 	return err
	// }
	// return viewTemplate.Execute(w, record)
	fmt.Println(r.URL.Path)
	return nil
}

//通过自定义路由器，简化处理异常的代码
func main() {
	http.Handle("/view", appHandler(viewRecord))
	http.ListenAndServe(":7777", nil)
}
