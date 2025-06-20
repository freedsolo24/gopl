package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func server1() {
	// /: 匹配所有的url
	// handler是处理函数
	// 当客户端访问路由是 "/" 的时候, 服务端会调用handler函数处理请求, handler的逻辑来如何响应
	// 服务器收到了 HTTP 的请求包，匹配了 HandleFunc 函数定义的 URL 路径，由 handler 函数做解析
	// 它是一个“默认兜底”的 handler, 它会匹配所有没有更具体匹配项的请求路径
	http.HandleFunc("/", handler)

	// 这里的hanler函数填入nil, 会使用底层全局默认的, 会和自己定义的handler形成呼应
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}

// 响应的来源于r.URL.Path
// handler 函数把 http 请求包进行解析，生成了 r 实例，这个实例包含了 HTTP 请求包的信息
// w形参是表面看ResponseWrite接口类，底层本质是response接口类实例，具有Write方法
func handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path: 客户端请求的 URL 中的「路径」部分（不包含域名、参数、锚点等）
	// 在服务器内部创建*Response结构类实例, 并注入到handler函数中
	// 调用fmt.Fprintf, 实际上是调用的w.Write(), 将后面的字符串"URL.Path="写入response结构体实例的缓冲区
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

// 这个函数的意思是当输入127.0.0.1:8000/ count++
// 当输入127.0.0.1:8000/count 之后可以看到曾经有多少个连接
func server2() {
	// 服务器为每一个匹配到的客户端连接都会开启一个协程,
	// 这样服务器可以并行处理多个连接请求
	http.HandleFunc("/", handlerMutex)
	http.HandleFunc("/count", counterMutex)
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err)
}

// 加了互斥锁
// 如果并发下, 两个连接同时更改count,会被不正确的修改
// 必须保证变量每次修改只有一个goroutine可以改
// 首先尝试获得锁, 获得锁之后加锁.并继续执行下面的代码
// 如果其他协程同时也尝试执行 handlerMutex 函数，它们会在 mu.Lock() 那一行阻塞，等待锁被释放。
// 这意味着在同一时刻，只有一个协程能执行到 mu.Lock() 后的代码（直到该协程释放锁）
// 当第一个协程执行到 mu.Unlock() 时，锁被释放，其他被阻塞的协程才能依次获得锁并继续执行
func handlerMutex(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL, Path=%q\n", r.URL.Path)
}

func counterMutex(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func server3() {
	http.HandleFunc("/", handlerServer3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// http.Request是一个结构体，里面是http请求
func handlerServer3(w http.ResponseWriter, r *http.Request) {
	// 打印http行首
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 遍历http头
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	// 取请求报文中的主机字段，打印服务端的主机
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	// 打印客户端的ip
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)

	// ParseForm: 解析表单
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// 遍历表单
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%s\n", k, v)
	}

	// os.Stdout, io.Discard, reponse writer 3个结构类,既是又是, 可以赋值给writer
}
