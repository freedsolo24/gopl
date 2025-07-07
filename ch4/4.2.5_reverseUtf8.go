// 修改reverse函数用于原地反转UTF­8编码的[]byte。是否可以不用分配额外的内存？
// 目标是“反转字符顺序”，不是反转单个字节: 世a界 -> 界a世
// 注意: 左右两边的文字字符长度可能不一样
package main

import (
	"unicode/utf8"
)

// 我自己的思路:
// (1) 循环, 一个从最左边循环, 一个从最右边循环
// (2) 拿最左边的文字字符的leftSize; 拿最右边文字字符的rightSize
// (3) 建立一个左边的临时缓存, len是左侧文字字符的leftSize; 建立一个右边的临时缓存, len是右侧文字字符的rightSize
// (4) 把左侧的文字字符(起始索引=0, 结尾索引=0+leftSize)拷贝到左边临时缓存, 把右侧的文字字符(起始索引=j-rightSize, 结尾索引=j)拷贝到右边临时缓存
// (5) 把右边临时缓存拷贝到左侧的位置, 长度是右边文字字符的rightSize(起始索引=0, 结尾索引=0+rightSize)
// (6) 把左边临时缓存拷贝到右侧的位置, 长度是左边文字字符的leftSize(起始索引=j-leftSize, 结尾索引=j)
// (7) 自增和自减两个循环变量
func reverseUtf1(b []byte) {
	// _, tailSize := utf8.DecodeLastRune(b)
	startIdx := 0
	endIdx := len(b)

	for i, j := startIdx, endIdx; i < j; {
		_, hSize := utf8.DecodeRune(b[i:])
		_, tSize := utf8.DecodeLastRune(b[:j])

		// 把两个文字字符占据的字节片段交换
		// 左边的文字字符: b[i:i+hSize]
		// 右边的文字字符: b[j:j+tSize]
		// 把这两个变长字节的位置交换, 需要用到中间变量
		// 两个 UTF-8 字符可能是不同长度，例如：
		// 左边是英文 A（1 字节）
		// 右边是汉字 你（3 字节）
		// 所以不能用简单地 byte-by-byte 互换，只能使用临时缓冲区保存一个

		// 新建两个临时变量, 把左右两个文字字符拷贝到临时变量中
		left := make([]byte, hSize)
		copy(left, b[i:i+hSize])
		right := make([]byte, tSize)
		copy(right, b[j-tSize:j]) // 最后一个文字字符起始索引[len(b)-它的size:一直到最后]

		// 将右边的文字字符复制到左边的位置, 因为左右两个文字字符不一样, 所以左边复制要用右边文字字符的tSize
		copy(b[i:i+tSize], right)

		// 将左边的文字字符复制到右边的位置, 因为左右两个文字字符不一样, 所以右边复制要用左边文字字符的hSize
		copy(b[j-hSize:j], left) // 把左边的拷贝到右边, 右边的起始位置[len(b)-左边的size:一直到最后]

		i += hSize
		j -= tSize
	}
}

// 思路:
// (1) 构建position变量, 里面存储每一个文字字符的起始索引.[ 0 3 6 9 ...]
// (2) 构建for循环, 两个变量i,j, i控制左边文字字符的起始索引i, j控制右边文字字符的起始索引len(position)-1
// (3) 左边文字字符的起始索引position[i], 结尾索引调用抽出来的函数position[i]+leftSize
// (4) 右边文字字符的起始索引position[j], 结尾索引调用抽出来的函数position[j]+rightSize
// (5) 交换左右两个文字字符, 构建一个临时变量, 做到两个文字字符的交换

func reverseUtf2(b []byte) {
	// position变量, 把每一个文字字符的起始位置索引, 添加到int类型的切片
	var positions []int
	for i := 0; i < len(b); {
		positions = append(positions, i)
		_, size := utf8.DecodeRune(b[i:])
		i += size
	}

	// 两两交换字符（按位置倒序）
	for i, j := 0, len(positions)-1; i < j; i, j = i+1, j-1 {
		// 左边文字字符的起始索引
		start1 := positions[i]
		// 右边文字字符的结尾索引
		end1 := nextRuneEnd(b, start1)

		// 左边文字字符的起始索引
		start2 := positions[j]
		// 右边文字字符的结尾索引
		end2 := nextRuneEnd(b, start2)

		// 交换两个字符的字节片段
		swapBytes(b, start1, end1, start2, end2)
	}
}

// 计算一个 UTF-8 字符的结束位置
func nextRuneEnd(b []byte, start int) int {
	_, size := utf8.DecodeRune(b[start:])
	return start + size
}

// 原地交换两个 UTF-8 字符的字节片段
func swapBytes(b []byte, start1, end1, start2, end2 int) {
	// 构建一个临时变量, 把左边的文字字符拷贝到临时变量中
	tmp := make([]byte, end1-start1)
	copy(tmp, b[start1:end1])

	// 右边文字字符拷到左边
	copy(b[start1:end1], b[start2:end2])

	// 左边文字字符拷到右边
	copy(b[start2:end2], tmp)
}

// 双指针思路和我的思路一样, 把左右两个文字字符的交换抽出来成为一个函数
// (1) 构建一个循环, 双变量i, j
// (2) 拿最左边的文字字符[left, leftSize]; 拿最右边文字字符[right,rightSize]. 到这一步和我的思路一样
// (3) 可以把swap文字字符抽出来成为一个函数
func reverseUtf3(b []byte) {
	for i, j := 0, len(b); i < j; {
		// decode 左边的文字字符, 从头往后找
		r1, size1 := utf8.DecodeRune(b[i:])
		if r1 == utf8.RuneError && size1 == 1 {
			// 非法字符
			break
		}

		// decode 右边的文字字符, 从末尾往前找
		r2, size2 := utf8.DecodeLastRune(b[:j])
		if r2 == utf8.RuneError && size2 == 1 {
			// 非法字符
			break
		}

		// 如果 i >= j - size2，说明已经对撞，停止
		if i >= j-size2 {
			break
		}

		// 左边文字字符是 [i:i+size1], 右边文字字符是 [j-size2:j]
		// i: 左侧文字字符的开始索引, size1: 左侧文字字符的大小
		// j-size2: 右侧文字字符的开始索引, size2: 右侧文字字符的大小
		swapUTF8(b, i, size1, j-size2, size2)

		// 更新指针
		i += size1
		j -= size2
	}
}

// i1, size1是左边文字字符的起始位置, size大小; i2, size2是右边文字字符的起始位置, size大小
func swapUTF8(b []byte, i1, size1, i2, size2 int) {
	// 把左边的文字字符拷贝到临时变量里
	tmp := make([]byte, size1)
	copy(tmp, b[i1:i1+size1])

	// 左边位置, size是右边文字字符的size <- 右边文字字符
	copy(b[i1:i1+size2], b[i2:i2+size2])

	// 右边位置, size是左边文字字符的size <- 左边文字字符
	copy(b[i2:i2+size1], tmp)
}
