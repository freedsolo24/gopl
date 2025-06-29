// 判断两个字符串是否同文异构, 如: silent和listen, hello和hlelo 是同文异构
package main

import (
	"sort"
	"strings"
)

// 我的思路:
// 两个字符串的长度不一样, 返回false
// 两个字符串长度一样, 两个字符串不相等, 返回false
// 两个字符串长度一样, 但是字符串不相等, 就要判断是否同源异构
// 构造一个bitmap:=map[int]int, flag
// (1) 遍历第一个字符串, 每次拿一个字符, 嵌套遍历第二个字符串, 每次拿一个字符, 这两个字符进行比较
// (2) 判断当map[idx]=1, 说明比较过了, 不能重复比较
// (3) 如果s2的这个字符之前没有比较过, 当c1==c2, 将flag++, 且bitmap=1, 以后就不重复比较
// (4) 注意continue和break的使用

func anagram1(s1, s2 string) bool {

	switch {
	case len(s1) != len(s2):
		return false
	case s1 == s2:
		return true
	}
	l := len(s1)
	bitmap := make(map[int]int, l)
	flag := 0

	for _, c1 := range s1 {
		for j, c2 := range s2 {

			if bitmap[j] == 1 {
				continue
			}
			if c1 == c2 {
				bitmap[j] = 1
				flag++
				break
			}
		}
	}
	if flag == l {
		return true
	} else {
		return false
	}

}

// 思路
// (1) 构造两个map, 分别统计两个字符串中每个字符出现的次数
// (2) 比较两个map: 遍历map1, 拿他的key, 在map2里面看同样的key, v值是否一样
func anagram2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	flag := 0

	for _, k := range s1 {
		m1[k]++
	}
	for _, k := range s2 {
		m2[k]++
	}
	for k, v := range m1 {
		if m2[k] == v {
			flag++
		}
	}
	if flag == len(m1) {
		return true
	} else {
		return false
	}

}

// 思路:
// (1) 把s1，s2 这两个string, 分隔, 放在[]string
// (2) sort包对[]string进行排序
// (3) 比较这两个[]string, 相同就标志位++
// (4) 最后判断标志位和长度是否一致
func anagram3(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	a := strings.Split(s1, "")
	b := strings.Split(s2, "")
	flag := 0

	sort.Strings(a)
	sort.Strings(b)

	for i := range a {
		if a[i] == b[i] {
			flag++
		}

	}
	if flag == len(s1) {
		return true
	} else {
		return false
	}

}
