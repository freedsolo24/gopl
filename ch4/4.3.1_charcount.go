// 统计从stdin中输入一篇文章, 统计这篇文章rune文字字符各有多少个, utf8从1个字节到4个字节各有多少个
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func charCount1() {
	// key是文字字符, value是int, 构建了底层数据结构, !=nil
	counts := make(map[rune]int)

	// const UTFMax=4, 代表utf8编码长度是4个字节, 比如部分emoji
	// var utflen[5]int 变成了数组, len=5的数组, 索引0~4, 需要用的是utflen[1]-utflen[4]代表1~4个字节, 所以要加1
	// utflen[1]表示1个字节, 是acsii码; utflen[3]表示3个字节, 是汉字; utflen[4]表示4个字节, 是emoji
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	// 创建一个带缓冲区的 Reader，用来高效读取标准输入（键盘输入）
	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		// 读到结尾退出, 在新行输入 ctrl+d 是EOF
		if err == io.EOF {
			break
		}
		// 如果出现其他错误（不是 EOF），输出错误并退出程序
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		// 如果解码出来的是 Unicode 的替代字符 � '\uFFFD'，且长度是 1 字节，说明这个字符非法，记为 invalid，跳过
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// 使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别
// 扩展为统计Unicode 文字字符的类别
// 字母（unicode.IsLetter）, 数字（unicode.IsDigit）, 标点（unicode.IsPunct）, 空白字符（unicode.IsSpace）, 控制字符（unicode.IsControl）
func charCount2() {
	// 构建map实例
	categories := make(map[string]int)
	invalid := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		r, n, err := reader.ReadRune()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "characters count:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			categories["letter"]++
		case unicode.IsDigit(r):
			categories["digit"]++
		case unicode.IsPunct(r):
			categories["punct"]++
		case unicode.IsSpace(r):
			categories["space"]++
		case unicode.IsControl(r):
			categories["control"]++
		default:
			categories["other"]++
		}
	}
	fmt.Println("Character category counts:")

	for k, v := range categories {
		fmt.Printf("%-8s:%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("invalid characters: %d\n", invalid)
	}
}
