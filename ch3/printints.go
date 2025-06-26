package main

import (
	"bytes"
	"fmt"
)

func intsToStrings(v []int) string {
	var buf bytes.Buffer
	// buf实例调用WriteByte方法, 往buf实例后面追加一个字节
	buf.WriteByte('[')
	for i, v := range v {
		if i > 0 {
			// buf实例调用WriteString方法, 往buf实例后面追加字符串
			buf.WriteString(", ")
		}
		// Fprintf要求第一个形参实现io.Writer接口, bytes.Buffer结构类实现了Write方法, 就实现了Writer接口, 可以赋值
		fmt.Fprintf(&buf, "%d", v)
		// 也可以写成这个: buf.WriteString((strconv.Itoa(v)))
	}
	buf.WriteByte(']')
	return buf.String()
}
