package main

import (
	"fmt"
	"unicode/utf8"
)

func unicodeutf8() {
	s := "Hello, 世界"
	// len函数打印字符串有几个字节
	fmt.Println(len(s))
	// RuneCountInString函数打印字符串有几个文字字符
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// 上面的函数可以用for range代替
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}
