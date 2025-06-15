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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server2() {
	http.HandleFunc("/", handlerMutex)
	http.HandleFunc("/count", counterMutex)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server3() {
	http.HandleFunc("/", handlerServer3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// w指明响应写到哪, r是请求
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}

// 加了互斥锁
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

func handlerServer3(w http.ResponseWriter, r *http.Request) {
	// 打印http行首
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 遍历http头
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
	// 打印服务端的主机
	fmt.Fprintf(w, "Host=%q\n", r.Host)
	// 打印客户端的ip
	fmt.Fprintf(w, "RemoteAddr=%q\n", r.RemoteAddr)

	// ParseForm: 解析表单
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// 遍历表单
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}

	// os.Stdout, io.Discard, reponse writer 3个结构类,既是又是, 可以赋值给writer
}
