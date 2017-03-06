package main

type Args struct {
	A, B int
}

func Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	args := Args{1, 2}
	var r *int
	//这里报错，不能这么用，要用Args的指针，这也是一个比较奇怪的地方，为什么RPC注册的服务的第一个参数是 args *Args 但是调用是传入的是args（这个args创建方式是如上这种，返回的是值类型）
	Multiply(args, r)
}
