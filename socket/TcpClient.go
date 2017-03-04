package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println(os.Stderr, "usage : %s host:port", os.Args[0])
	// 	os.Exit(1)
	// }
	service := "127.0.0.1:7777"

	//创建一个tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//开启一个tcp连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	//写入数据
	// _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	_, err = conn.Write([]byte("timestamp"))
	// for i := 0; i < 3; i++ {
	// 	conn.Write([]byte("abc"))
	// }
	// conn.Write([]byte("timestamp"))
	// conn.Write([]byte("timestamp"))
	checkError(err)

	//拿到返回的数据,服务端响应反馈的信息
	//这里注意一下！！！这个地方不知道为啥，是在整个链接结束之后才从conn中拿出数据，不是服务端一边写一边拿，很奇怪
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(result)
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
