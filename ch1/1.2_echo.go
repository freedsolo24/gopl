package main

import (
	"fmt"
	"os"
	"strings"
)

// 每次保存自动运行gofmt
// os.Args 是os包的Args变量名, 类型[]string, 里面的元素是命令行的参数
// os.Args[0]=命令的绝对路径, os.Args[1]=参数1, os.Args[2]=参数2, ...

// 调试输入: go run . 1 2 3
// 调试输出: Index[0], Value=/tmp/go-build3498032286/b001/exe/ch1
//
//	Index[1], Value=1
//	Index[2], Value=2
//	Index[3], Value=3
func echo0() {
	for i, str := range os.Args {
		fmt.Printf("Index[%d], Value=%s\n", i, str)
	}
}

// echo1(for c版): 把所有的参数打印到一行, 把空格写前面, 不会在最后拼接完末尾多一个空格的情况
func echo1() {
	line, sep := "", ""

	for i := 0; i < len(os.Args); i++ {
		if i >= 1 {
			line = line + sep + os.Args[i]
			sep = " "
		}
	}
	fmt.Printf("echo1  行: %s,长度: %d\n", line, len(line))
}

// echo2(for range版): 把所有的参数打印到一行
func echo2() {
	line, sep := "", ""

	for _, str := range os.Args[1:] {
		// 把空格写前面, 不会产生最后拼接完多一个空格的情况
		line += sep + str
		sep = " "
	}
	fmt.Printf("echo2  行: %s,长度: %d\n", line, len(line))
}

// echo3(strings.Join 作用: 把[]string切片, 按照指定分隔符拼接成一行)
func echo3() {
	// Strings.Join函数, 要求第一个形参类型[]string,第二个形参是分隔符
	// 第一个形参用os.Args的子切片从Index=1开始

	line := strings.Join(os.Args[1:], " ")

	fmt.Printf("echo3  行: %s,长度: %d\n", line, len(line))
}
