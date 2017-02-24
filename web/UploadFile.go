package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe("127.0.0.1:9998", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func upload(w http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method) //获取请求的方法
	if request.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("UploadFile.gtpl")
		t.Execute(w, token)
	} else {
		request.ParseMultipartForm(32 << 20)
		file, handler, err := request.FormFile("uploadfile")
		defer fmt.Println("第一个defer执行了")
		defer file.Close()
		if err != nil {
			fmt.Println("request.formFile error")
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%v", handler.Header)
		fmt.Println("~/" + handler.Filename)
		f, err := os.OpenFile("./"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		// defer fmt.Println("第二个defer执行了")
		// defer f.Close()     <-- 这个defer也是执行的，但是放在下面大概是习惯问题
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fmt.Println("第二个defer执行了")
		defer f.Close()
		io.Copy(f, file)
	}
}

//文件handler是multipart.FileHeader,里面存储了如下结构信息
//type FileHeader struct {
// Filename string
// Header   textproto.MIMEHeader
// contains filtered or unexported fields
// }
