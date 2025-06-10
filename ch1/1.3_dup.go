package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打印重复的行, 输出stdin中出现2次及以上的行,涉及 map, bufio包. 以下案例是流式输入.

// NewScanner函数, 初始化一个scanner, 对stdin做逐行读取器的实例句柄
// scanner.Scan()：尝试读取下一行输入，成功则返回 true
// scanner.Text()：返回刚刚读入的一行字符串（去掉换行）
// 从stdin读取输入,把行放到key中,把次数放入value中. 最后打印value中大于1的key值
func dup1() {
	// 构建一个map容器实例, key=行, value=次数
	counts := make(map[string]int)
	// 初始化一个scanner, 对系统的stdin buffer进行扫描的工具, 相当于初始化一个对stdin buffer操作的句柄
	scanner := bufio.NewScanner(os.Stdin)

	// 阻塞在这里, 调用Scan方法, 从stdin输入流中尝试读取一行(我输一行他读一行), 读到自己的临时buffer
	// 直到我按ctrl+d, Scan返回false, 循环结束
	for scanner.Scan() {
		// 调用Text方法, 把临时buffer中的内容取出, 是一个string, 放到map[key]中, value+1
		counts[scanner.Text()] = counts[scanner.Text()] + 1
	}
	// 最后遍历, value>1打印对应的key
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("重复的行: %s, 重复次数: %d\n", line, n)
		}
	}
}

// 从文件列表或标准输入读取
func dup2() {
	// make一个容器
	counts := make(map[string]int)
	// 取到files, files是一个[]string
	files := os.Args[1:]
	// 如果files长度=0, 说明没有参数, 也就是参数, 那么就读取stdin, 塞到counts里
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("stdin 重复的行:%s, 重复的次数:%d\n", line, n)
			}
		}
		// 隐含条件有参数, >=1个
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			fmt.Printf("文件:%s\n", file)
			countLines(f, counts)
			for str, n := range counts {
				if n > 1 {
					fmt.Printf("重复的行:%s, 重复的次数:%d\n", str, n)
				}
			}
			fmt.Println()
			clear(counts)
		}
	}
}

func dup3() {

}

func countLines(r *os.File, m map[string]int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		m[scanner.Text()] = m[scanner.Text()] + 1
	}
}
