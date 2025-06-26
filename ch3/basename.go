// 注意:(1) 循环从后往前 (2) 字符串里面的字符匹配要用单引号 (3) 匹配之后退出循环
//
//	(4) 从后往前循环字符串的当前位置, 往后截取, 通过索引截
//	(5) 从后往前循环字符串的当前位置, 往前截取, 通过容量截
package main

import "strings"

// a=>a, a.go=>a, a/b/c.go=>c, a/b.c.go=>b.c
func basename1(s string) string {
	// 循环从后往前判断, 通过索引取当前索引到末尾的字符串
	for i := len(s) - 1; i >= 0; i-- {
		// 用单引号匹配, 因为s是string类型, 里面的一个字符是byte类型, 单引号表示里面是一个rune类型,或byte类型
		// Go允许byte和rune比较, 因为Go做了隐式类型转换 int32(s[i]) == '/'
		if s[i] == '/' {
			s = s[i+1:]
			// 截取之后就要退出循环
			break
		}
	}
	// 循环从后往前判断, 通过容量, 从头到中间截取
	// a.b.jpg   一个7个元素 索引0-6 从后往前删 654都不要 索引3是. 前面索引是[0-2]一共3个元素都要
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			// 截取之后退出循环
			break
		}
	}
	return s
}

// a=>a, a.go=>a, a/b/c.go=>c, a/b.c.go=>b.c   b.c.go
func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	dot := strings.LastIndex(s, ".")
	// dot=3, s[3]=. 说明前面有3个元素, 索引是s[0] s[1] s[2], 容量就写dot
	s = s[:dot]
	return s
}
