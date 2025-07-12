package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// // sum1, sum0 := convshaPopcount("abc")
	// // fmt.Printf("字符串哈希后的1和0各有: 1有%d个bit, 0有%d个bit\n", sum1, sum0)
	// // fmt.Printf("两个字符串哈希值不同的bit数:%d\n", diffbit("abc", "abcd"))
	// flagstdin()
	// flagAlgo()
	// a := [...]int{0, 1, 2, 3, 4, 5}
	// l := [...]int{1, 2, 3, 4, 5, 6, 7}
	// r := [...]int{1, 2, 3, 4, 5, 6, 7}
	// // a[:] 将数组a转换成切片类型 []int  a[:]等价于 []int{0,1,2,3,4,5}
	// // 因为传进去的是切片, 所以修改会共用底层数组a
	// // a[0:len(s)]=a[0:6] 意思是从0开始拿6个元素 0 1 2 3 4 5 结尾的索引到5
	// reverseSlice(a[:])
	// fmt.Println("就地反转", a)

	// // 将 slice 向左旋转 n 个元素，意味着把开头的前 n 个元素搬到 slice 的末尾，其他元素前移。
	// // 原始切片: [1 2 3 4 5 6 7]  左旋3位: [4 5 6 7 1 2 3]
	// // 原始顺序：A B  →  期望结果：B A （左旋）
	// // 第一次反转 A → A 顺序被打乱
	// // 第二次反转 B → B 顺序也被打乱
	// // 第三次整体再反转 → 两部分的“打乱顺序”叠加后，恰好成了正确的新顺序！
	// reverseLeft(l[:2]) // 反转前 n 个元素       // [3 2 1 4 5 6 7]
	// reverseLeft(l[2:]) // 反转后 len-n 个元素   // [3 2 1 7 6 5 4]
	// reverseLeft(l[:])  // 反转整个切片          // [4 5 6 7 1 2 3]
	// fmt.Println("就地左旋2次 ", l)

	// // 将 slice 向右旋转 n 个元素，意味着把末尾的后 n 个元素搬到 slice 的开头，其他元素后移。
	// // 原始切片: [1 2 3 4 5 6 7]  右旋3位: [5 6 7 1 2 3 4]
	// // 原始顺序：B A  →  期望结果：A B （右旋）
	// reverseRight(r[4:]) // 反转后 n 个元素       [1 2 3 4 7 6 5]
	// reverseRight(r[:4]) // 反转前 len-n 个元素   [4 3 2 1 7 6 5]
	// reverseRight(r[:])  // 整体反转              [5 6 7 1 2 3 4]
	// fmt.Println("就地右旋3次 ", r)
	// // 原始： A B      → [1 2 3 4 | 5 6 7]
	// // Step1: A B'     → [1 2 3 4 | 7 6 5]
	// // Step2: A' B'    → [4 3 2 1 | 7 6 5]
	// // Step3: (A' B')' → [5 6 7 | 1 2 3 4]

	// array := [6]int{1, 2, 3, 4, 5, 6}
	// reverseArray(&array)
	// fmt.Println("反转数组 ", array)

	// s := []int{1, 2, 3, 4, 5, 6, 7}
	// // rotateLeft(s, 3)
	// rotateRight(s, 3)
	// fmt.Println("一次循环 ", s)

	// s1 := []string{"a", "a", "b", "b", "b", "c", "a", "a"}
	// fmt.Println(chRepeat2(s1))

	// b1 := []byte("    Hello      world!     ")
	// b1 = squashSpace1(b1)
	// fmt.Printf("b1压缩空格|%s\n", string(b1))

	// b2 := []byte("   Hello   world!    |")
	// b2 = squashSpace2(b2)
	// fmt.Printf("b2压缩空格|%s\n", string(b2))

	// input := []byte("你好　  世界\t\t\t！")
	// result := squashUnicodeSpaces1(input)
	// fmt.Printf("结果: %q\n", result)

	// b3 := []byte("世￡ム界aプb�你好￥")
	// reverseUtf2(b3)
	// fmt.Printf("反转：%s\n", string(b3)) // 界a世

	// // charCount2()

	// jsonMarshal()
	// result, err := SearchIssues1(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	// r, err := SearchIssues2(os.Args[1:])
	// if err != nil {
	// 	panic(err)
	// }
	// categories(r)
	// xkcd1()

	// 定义子命令
	downloadCmd := flag.NewFlagSet("download", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)

	// 定义搜索的关键词
	searchTerms := searchCmd.String("terms", "", "Search terms (comma-separated)")

	// 命令行参数小于2, 提出使用说明
	if len(os.Args) < 2 {
		fmt.Println("Usage: xkcd-tool <Subcommand> [options]")
		fmt.Println("Subcommands: download, search")
		os.Exit(1)
	}

	// os.Args[0]=xkcd-tool, os.Args[1]=download | search
	switch os.Args[1] {
	case "download":
		// os.Args[2:] 指的是什么
		downloadCmd.Parse(os.Args[2:])
		if err := downloadComics(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "search":
		searchCmd.Parse(os.Args[2:])
		if *searchTerms == "" {
			searchCmd.Usage()
			os.Exit(1)
		}
		terms := strings.Split(strings.ToLower(*searchTerms), ",")
		if err := searchComics(terms); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Command: download, search")
		os.Exit(1)
	}

}

// go run . repo:golang/go is:open sort:created-aesc json decoder
