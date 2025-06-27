// 这个函数截取一段字符串的数字, 从后往前, 3位截成一段插入一个逗号
// 位于中间, 后面的全都要, 用索引, s[n:]
// 位于中间, 前面的全都要, 用容量, s[:n]
package main

import (
	"bytes"
	"strings"
)

// 12345=>12,345
// 如果是3个元素就返回, 如果是4个元素, 后3个元素不要, 只要1个, 123456, 12345678
// 123456789=>123456
// 123456=>123
// 123,456
func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	// 从头到中间截取, 通过容量, s[:n-3]
	// 前面的不要, 从中间要到末尾, 通过索引
	// comma是一个递归函数,返回段是return 3
	return comma1(s[:n-3]) + "," + s[n-3:]
}

// 递归函数是如何“分解 + 合并”的过程: 如果输入123456789, 是怎么递归变成123,456,789

//                                                              (7) 最终函数返回 "123,456,789"
// (1) comma("123456789") n=9>3 ->继续递归
//     return comma(s[:6]),+,s[6:]         comma(123456)+,+789  (6) return "123,456"+","+789  向上返回 "123,456,789"
// (2) comma("123456")    n=6>3 ->继续递归
//     return comma(s[:3]),+,s[3:]         comma(123)+,+456     (5) return "123"+","+456      向上返回 "123,456"
// (3) comma("123")       n=3   ->终止递归
// (4) return "123" 向上返回

// comma("123456789")
// ↳ comma("123456") + "," + "789"
//    ↳ comma("123") + "," + "456"
//       ↳ "123"
//    ← "123,456"
// ← "123,456,789"

// 1234567
// 7 c=1  76 c=2  765 c=3 , 765,4
// 1,234,567

// 765,432,1

// 我的思路: 反着拿字符串, 拿3个,插一个逗号, 在拿三个, 得到逆序的字符串
//
//	把逆序的字符串, 在反着拿, 放到buffer里面, 就变成了正序
func comma2(s string) string {
	var buf bytes.Buffer
	length := len(s)
	c := 0
	for idx := length - 1; idx >= 0; idx-- {
		if c == 3 {
			buf.WriteByte(',')
			c = 0
		}
		buf.WriteByte(s[idx])
		c++
	}

	// reverse := buf.String()
	// var buf2 bytes.Buffer

	// for i := len(reverse) - 1; i >= 0; i-- {
	// 	buf2.WriteByte(reverse[i])
	// }
	// return buf2.String()

	// 输出是一个[]byte, 可以基于他就地逆序, 如果是buf.String, 不能在字符串的基础上做操作
	b := buf.Bytes()

	reverseBytes(b)
	return string(b)
}
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// 正着拿的思路: 利用余数, 长度处于3, 取余数, 余几个就是甩几个, 余0, 就是正好3的倍数
//
//	因为多余的在前面, 所以先把多余的截取出, 利用子切片
//	剩下的开始循环, 先插入逗号, 在3个一取插入, 后面的都是3的倍数
func comma3(s string) string {
	var buf bytes.Buffer
	n := len(s)

	// 前面余数位数不带逗号, 4%3余1,pre=1 5%3余2,pre=2  7%3余1,pre=1 8%3余2,pre=2
	pre := n % 3
	// 6%3余0,pre=3 9%3余0,pre=3
	if pre == 0 && n >= 3 {
		pre = 3
	}

	// 先写前几位, 相当于把余数的部分写了, 后面都是3的倍数了
	buf.WriteString(s[:pre])

	// 后面都是3个倍数了, 3个一组加逗号
	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}

// 处理小数点和负号,123456->123,456 -123456->-123,456 -1234567.89->-1,234,567.89
// 思路: (1) 利用strings.HasPrefix函数,判断字符串的前缀 (2) 利用string.SplitN, 把sting以.分成两块放到[]string (3) 整数部分放在一个变量里, 小数部分放在一个变量里
// (4) 取出整数部分之后, 给整数部分加逗号 (5)拼接小数部分  (6)拼接符号
func comma4(s string) string {
	// 处理负号
	negative := false
	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	}

	// 分离整数部分和小数部分
	var intPart, fracPart string
	parts := strings.SplitN(s, ".", 2)
	intPart = parts[0]
	if len(parts) == 2 {
		fracPart = parts[1]
	}

	// 给整数部分加逗号
	var buf bytes.Buffer
	n := len(intPart)
	pre := n % 3
	if pre == 0 && n >= 3 {
		pre = 3
	}
	buf.WriteString(intPart[:pre])
	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(intPart[i : i+3])
	}

	// 拼接小数部分
	if fracPart != "" {
		buf.WriteByte('.')
		buf.WriteString(fracPart)
	}

	// 拼接负号
	result := buf.String()
	if negative {
		result = "-" + result
	}

	return result
}
