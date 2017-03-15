package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}

/*
注意以下几点
	传递参数-ldflags "-s"，忽略debug的打印信息
	传递-gcflags "-N -l" 参数，这样可以忽略Go内部做的一些优化，聚合变量和函数等优化，这样对于GDB调试来说非常困难，所以在编译的时候加入这两个参数避免这些优化。
*/

/**
gdb常用指令
	list
	简写命令l，用来显示源代码，默认显示十行代码，后面可以带上参数显示的具体行，例如：list 15，显示十行代码，其中第15行在显示的十行里面的中间，如下所示。

	break
	简写命令 b,用来设置断点，后面跟上参数设置断点的行数，例如b 10在第十行设置断点。

	delete
	简写命令 d,用来删除断点，后面跟上断点设置的序号，这个序号可以通过info breakpoints获取相应的设置的断点序号.

	backtrace
	简写命令 bt,用来打印执行的代码过程

	info
	info命令用来显示信息，后面有几种参数

	print
	简写命令p，用来打印变量或者其他信息，后面跟上需要打印的变量名，当然还有一些很有用的函数$len()和$cap()，用来返回当前string、slices或者maps的长度和容量。

	whatis
	用来显示当前变量的类型，后面跟上变量名，例如whatis msg

	next
	简写命令 n,用来单步调试，跳到下一步，当有断点之后，可以输入n跳转到下一步继续执行

	coutinue
	简称命令 c，用来跳出当前断点处，后面可以跟参数N，跳过多少次断点

	set variable
	该命令用来改变运行过程中的变量值，格式如：set variable <var>=<value>
*/
