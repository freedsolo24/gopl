// 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

// 1. 按单词读取输入, Go 默认 scanner.Scan() 是按“行”来读取的，这里显式设置为 按单词读取 scanner.Split(bufio.ScanWords)
// 2. 声明 map 存放 word => count. 这个 map 用来统计：每个单词出现了多少次
// 3. 循环读取并统计
// NewScanner实现
func wordFreq1() {

	wFreq := make(map[string]int)

	// 生成对标准输入的扫描器句柄
	scanner := bufio.NewScanner(os.Stdin)
	// 默认是行读取, 这里这句话的意思是告诉scanner用"单词"的方式来分割输入
	scanner.Split(bufio.ScanWords)
	fmt.Println("请输入文本, ctrl+d 结束")

	for scanner.Scan() {
		word := scanner.Text()
		wFreq[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取出错: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("words\tfreq\t")
	for k, v := range wFreq {
		fmt.Printf("%-8s, %d\n", k, v)
	}
}

// NewReader实现, 自己用 strings.Fields() 读取一行
// 自己用 strings.Split() 分词
func wordFreq2() {
	freq := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "read error:%v\n", err)
			break
		}
		// 对一行进行分词
		words := strings.Fields(line)

		for _, word := range words {
			freq[word]++
		}
		if err == io.EOF {
			break
		}
	}
	for k, v := range freq {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func wordFreq3() {
	// 使用 Scanner 逐词读取输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// 词频 map
	freq := make(map[string]int)

	// 用正则表达式清除词两边的标点
	re := regexp.MustCompile(`^[\pP\s]+|[\pP\s]+$`) // 清除前后标点（中文英文标点都能去）

	for scanner.Scan() {
		word := scanner.Text()

		// 转小写，去除标点
		word = strings.ToLower(word)
		word = re.ReplaceAllString(word, "")

		if word != "" {
			freq[word]++
		}
	}

	// 收集单词并排序
	type pair struct {
		word  string
		count int
	}

	var pairs []pair
	totalWords := 0
	for word, count := range freq {
		pairs = append(pairs, pair{word, count})
		totalWords += count
	}

	// 按频率降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// 打印结果
	fmt.Println("单词频率：")
	for _, p := range pairs {
		fmt.Printf("%-15s: %d\n", p.word, p.count)
	}

	fmt.Printf("\n总单词数：%d\n", totalWords)
	fmt.Printf("不同单词数：%d\n", len(freq))
}
