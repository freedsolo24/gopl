package main

import "fmt"

func main() {
	sum1, sum0 := convshaPopcount("abc")
	fmt.Printf("字符串哈希后的1和0各有: 1有%d个bit, 0有%d个bit\n", sum1, sum0)
}
