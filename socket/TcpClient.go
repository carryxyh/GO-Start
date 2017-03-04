package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "usage : %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	//创建一个tcp地址
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	//开启一个tcp连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	//写入数据
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	//拿到返回的数据,服务端响应反馈的信息
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
