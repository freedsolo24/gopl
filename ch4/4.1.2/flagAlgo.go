package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func flagAlgo() {

	// algo类型是一个指针*string
	// algo指向原始的无命名变量标头值
	// flag.String内部自动生成了一个无名变量, 内容是用户输入的字符串
	var algo = flag.String("algo", "sha256", "hash algorithm: sha256, sha384, sha512")
	flag.Parse()

	// 拿非flag参数的第一个, 没有就返回""
	input := flag.Arg(0)

	if input == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			str := scanner.Text()
			sum := sha256.Sum256([]byte(str))
			fmt.Printf("SHA256: %x\n", sum)
		}
		// 一旦用户按下 Ctrl+D（Linux/macOS）表示 EOF（文件结束），
		// 此时 scanner.Scan() 返回 false，跳出 for 循环，执行 return
		//  如果输入流结束（EOF），它就会跳出.所以需要 return，确保 程序在标准输入处理完毕后结束，不继续往下执行 switch 分支。
		// 按 Ctrl+D 结束输入，你会发现程序优雅退出，不会进入 switch 分支
		return
	}

	// *algo拿的是原始的无名变量标头值, 是一个变量, 下面的case就是变量的各种值
	// string虽然有标头值, 本质还是值类型, 要用指针
	switch *algo {
	case "sha256":
		sum := sha256.Sum256([]byte(input))
		fmt.Printf("SHA256: %x\n", sum)
	case "sha384":
		sum := sha512.Sum384([]byte(input))
		fmt.Printf("SHA384: %x\n", sum)
	case "sha512":
		sum := sha512.Sum512([]byte(input))
		fmt.Printf("SHA512: %x\n", sum)
	default:
		fmt.Fprintf(os.Stderr, "Unsupported algorithm: %s\n", *algo)
		os.Exit(1)
	}

}

func flagstdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		sum := sha256.Sum256([]byte(str))
		fmt.Printf("sha256 %x\n", sum)
	}
}
