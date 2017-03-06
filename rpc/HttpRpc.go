package main

//Go RPC的函数只有符合下面的条件才能被远程访问，不然会被忽略，详细的要求如下：
// 函数必须是导出的(首字母大写)
// 必须有两个导出类型的参数，
// 第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
// 函数还要有一个返回值error

//func (t *T) MethodName(argType T1, replyType *T2) error
// * T、T1和T2类型必须能被encoding/gob包编解码。

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Arith int

type Quotient struct {
	Quo, Rem int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	//注册了一个Arith的RPC服务
	rpc.Register(arith)
	//rpc.HandleHTTP函数把该服务注册到了HTTP协议上，然后我们就可以利用http的方式来传递数据了
	rpc.HandleHTTP()

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println(err)
	}
}
