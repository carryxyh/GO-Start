### Go语言中的Dispatcher<br/>

Go使用goroutines来处理connection的读写事件，不会阻塞：<br/>

```
	c, err := srv.newConn(rw)
    if err != nil {
        continue
    }
    go c.serve()
```
<br/>

c即为创建的connection，保存了该次请求的信息，然后再传递到对应的handler，handler就可以读取到请求的header信息，保证了请求之间独立。<br/>

#### Go中的ServeMux<br/>

上面代码中提到了c（这个c就是connection）.serve（）方法。其实内部是调用了http包默认的路由器，通过路由器把本次请求的信息传递到了后端的处理函数。<br/>

默认路由器`ServeMux`，结构如下：<br/>

```
type ServeMux struct {
	mu sync.RWMutex   //锁，由于请求涉及到并发处理，因此这里需要一个锁机制
	m  map[string]muxEntry  // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
	hosts bool // 是否在任意的规则中带有host信息
}
```
<br/>

下面看一下`muxEntry`：<br/>

```
type muxEntry struct {
	explicit bool   // 是否精确匹配
	h        Handler // 这个路由表达式对应哪个handler
	pattern  string  //匹配字符串
}
```
<br/>

接着看一下`Handler`的定义：<br/>
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}
```
<br/>

`Handler`是一个接口，但是前一小节中的`sayhelloName`函数并没有实现ServeHTTP这个接口，仍然能添加到路由表中，原因就是http包里还有一个`HandlerFunc`，我们定义的函数`sayhelloName`就是这个HandlerFunc调用的结果，而这个类型默认实现了ServeHTTP这个接口，即我们调用了HandlerFunc（f），强制类型转换f成为`HandlerFunc`类型，这样f就拥有了ServeHTTP方法。<br/>

```
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
<br/>

我们看一下HandlerFunc的官方注解：<br/>
> HandlerFunc类型是一个适配器，允许使用普通的函数作为HTTP处理程序。如果f是具有适当签名的函数，HandlerFunc（f）是调用f的Handler。

<br/>
适当的签名，由于作者水平也不深厚（毕竟我本命语言是java），猜一下指的应该是函数的参数以及返回值，也就是说：***如果函数的参数是两个，分别是ResponseWriter和一个指向Request的指针，并且返回值为void类型的函数，可以强转为HandlerFunc，而最终调用的f中的Handler接口的方法也就是ServeHttp***。
<br/>

路由器里面存储好了相应的路由规则之后，那么具体的请求又是怎么分发的呢？请看下面的代码，默认的路由器实现了ServeHTTP：<br/>

```
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		w.Header().Set("Connection", "close")
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

如上所示路由器接收到请求之后，如果是`*`那么关闭链接，不然调用`mux.Handler(r)`返回对应设置路由的处理Handler，然后执行`h.ServeHTTP(w, r)`。看一下`ServeMUX.Handler（*request）`的官方文档：<br/>
> Handler返回用于给定请求的处理程序，请咨询r.Method，r.Host和r.URL.Path。它总是返回一个非nil处理程序。如果路径不是其规范形式，处理程序将是重定向到规范路径的内部生成的处理程序。
> Handler还返回与请求匹配的注册模式，或者在内部生成的重定向的情况下，返回在跟随重定向之后匹配的模式。
> 如果没有适用于请求的注册处理程序，则Handler返回“未找到页面”处理程序和空模式。
 
说白了，根据request的method、host和请求的URL的路径返回一个处理程序，这个处理程序就是我们说过的Handler，再看看Handler接口的方法，我们就知道了，最终会跑到我们`sayhelloName`里面~。我们看看`ServeMux.Handler(*request)`的实现：<br/>

```
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
	if r.Method != "CONNECT" {
		if p := cleanPath(r.URL.Path); p != r.URL.Path {
			_, pattern = mux.handler(r.Host, p)
			return RedirectHandler(p, StatusMovedPermanently), pattern
		}
	}   
	return mux.handler(r.Host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	mux.mu.RLock()
	defer mux.mu.RUnlock()

	// Host-specific pattern takes precedence over generic ones
	if mux.hosts {
		h, pattern = mux.match(host + path)
	}
	if h == nil {
		h, pattern = mux.match(path)
	}
	if h == nil {
		h, pattern = NotFoundHandler(), ""
	}
	return
}
```
<br/>

为了不让读者懵逼，我们还是看一下match方法，这是个私有方法，循环迭代了mux中的map：<br/>

```
func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}
```

匹配到之后返回存储的handler，调用这个handler的ServeHTTP接口就可以执行到相应的函数了。<br/>

Go其实支持外部实现的路由器 ListenAndServe的第二个参数就是用以配置外部路由器的，它是一个Handler接口，即外部路由器只要实现了Handler接口就可以,我们可以在自己实现的路由器的ServeHTTP里面实现自定义路由功能。<br/>

我们实现一个简易路由器：<br/>
 
```
package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
```

通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程：<br/>

* 首先调用Http.HandleFunc，按顺序做了几件事：<br/>

	1. 调用了DefaultServeMux的HandleFunc<br/>

	2. 调用了DefaultServeMux的Handle<br/>

	3. 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则<br/>

* 其次调用http.ListenAndServe(":9090", nil)，按顺序做了几件事情：<br/>

	1. 实例化Server<br/>

	2. 调用Server的ListenAndServe()<br/>

	3. 调用net.Listen("tcp", addr)监听端口<br/>

	4. 启动一个for循环，在循环体中Accept请求<br/>

	5. 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()<br/>

	6. 读取每个请求的内容w, err := c.readRequest()<br/>

	7. 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux<br/>

	8. 调用handler的ServeHttp<br/>

	9. 在这个例子中，下面就进入到DefaultServeMux.ServeHttp<br/>

	10. 根据request选择handler，并且进入到这个handler的ServeHTTP，`mux.handler(r).ServeHTTP(w, r)`<br/>

	11. 选择handler：<br/>

  * 判断是否有路由能满足这个request（循环遍历ServerMux的muxEntry）

 * 如果有路由满足，调用这个路由handler的ServeHttp

 * 如果没有路由满足，调用NotFoundHandler的ServeHttp

感谢作者，[原文地址](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.4.md)