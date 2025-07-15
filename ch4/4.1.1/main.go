// 统计sha256散列中不同的位数
package main

import "fmt"

func main() {
	sum1, sum0 := convshaPopcount("abc")
	fmt.Printf("字符串哈希后的1和0各有: 1有%d个bit, 0有%d个bit\n", sum1, sum0)
	fmt.Printf("两个字符串哈希值不同的bit数:%d\n", diffbit("abc", "abcd"))
}
