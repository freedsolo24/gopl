package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// 从指定的url里, 解析信息
// 这相当于写了一个客户端
func fetch1() {

	urls := os.Args[1:]
	for _, url := range urls {

		// Get函数对url解析, 拿到response结构体的实例, 要对response结构体实例做解析
		// 发起tcp连接, 发送对于url的http请求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}

		// resp.Body
		// 从resp结构类的body属性中读数据
		// resp.Body类型是ReadCloser
		// resp.Body并不会将http响应包中的body数据读到内存中, 而是由resp.body提供了从"远端服器流式读取的接口"
		// 这个接口背后连接着底层的tcp网络连接

		// resp.Body接口类实例的底层是http.body实例, 传入ReadAll函数, 最终是http.body.调用了Read方法
		b, err := io.ReadAll(resp.Body)

		// 用完之后, 连接通道关闭, 释放连接资源
		// 这里没有err, 是因为: 关闭失败通常不影响程序功能, 可以省略
		resp.Body.Close()

		// 这里的err是ReadAll的err
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}

func fetchCopy() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url %s %v\n", url, err)
			os.Exit(1)
		}

		defer resp.Body.Close()

		// :=的作用是声明新变量
		// err前面定义过,在这里就不是新变量. 如果用_,也不算是新变量
		// 所以就得用=来定义
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stdout, "%v\n", err)
		}

	}
}

func fetchPrefix() {
	urls := os.Args[1:]
	var body []byte
	for _, url := range urls {
		if !strings.HasPrefix(url, "https://") {
			url = fmt.Sprintf("https://%s", url)
		}
		body = respBody(url)
		fmt.Printf("url %s的body部分\n", url)
		fmt.Printf("%s\n", string(body))
		fmt.Println(strings.Repeat("-", 60))
	}
}

func fetchCode() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch url %s %v\n", url, err)
			os.Exit(1)
		}

		fmt.Println(resp.Status)
	}
}

func respBody(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch url %s %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch reading url body %s %v\n", url, err)
		os.Exit(1)
	}
	return body
}
