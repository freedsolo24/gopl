package main

import (
	"fmt"
)

func main() {
	// args := os.Args[1:]
	// for _, arg := range args {
	// 	fmt.Printf("origin string:%s\n", arg)
	// 	fmt.Printf("              %s\n", comma(arg))

	// }
	// fmt.Println(intsToStrings([]int{1, 2, 3, 4, 5, 6, 7}))

	// fmt.Println(comma4("-123456.89"))

	// // 实例v的标志位=Multicast+Up
	// // v的类型本质是uint8, 是值类型, 所以设置v, 要传指针
	// var v Flags = FlagMulticast | FlagUp

	// // 查看v是不是Up
	// fmt.Printf("%b %t\n", v, IsUp(v))
	// // 把标志位v设置down
	// TurnDown(&v)
	// // 把标志位v设置广播
	// SetBroadcast(&v)
	// fmt.Printf("%b %t\n", v, IsUp(v))
	// fmt.Printf("%b %t\n", v, IsCast(v))
	fmt.Println(anagram1("abcdc", "dccab"))

}
