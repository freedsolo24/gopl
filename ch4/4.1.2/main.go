// 在默认情况下输出其标准输入的sha256散列, 但也支持输出sha384或sha512散列的命令行标记
package main

func main() {
	flagstdin()
	flagAlgo()
}
