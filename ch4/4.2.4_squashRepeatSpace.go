// 在原地将一个 UTF-8 编码的 []byte 中，相邻的空白字符（空格、制表符、换行等）替换为一个空格。
// 使用两个索引：  readIdx：读的位置, writeIdx：写的位置
// 维护一个标志位：prevIsSpace：前一个 rune 是否是空白字符
package main

import (
	"unicode"
	"unicode/utf8"
)

// 仅能输入ascii码, 流程
// 1. 从左向右读字符（注意是 UTF-8，多字节的字符要用 utf8.DecodeRune 解析）
// 2. 判断当前字符是否是空白字符（用 unicode.IsSpace(rune)）
// 3. 如果当前是空格，并且前一个不是空格，则写入一个 ' ' 到 writeIdx
// 4. 如果当前是空格，并且前一个也是空格，则跳过（不写）
// 5. 如果当前不是空格，则正常写入字符，并更新 writeIdx
// 6. 最终返回 b[:writeIdx] 即为结果
// 原则: 只有当 "空格空格" 的时候才跳过, 不做复制, 不复制说明不要. 只有复制说明要这个位.
func squashSpace1(b []byte) []byte {

	writeIdx := 1

	for readIdx := 1; readIdx < len(b); {
		if b[readIdx] == ' ' && b[readIdx-1] == ' ' {
			readIdx++
			continue
		} else {
			b[writeIdx] = b[readIdx]
			readIdx++
			writeIdx++
		}
	}
	// 就是就地返回的, 必须返回一个b[:writeIdx], 如果没有用原字符串, 原字符串打印会是: Hello world!  !d
	return b[:writeIdx]
}

// 仅能输入ascii码,
// 使用append, 删除重复的"空格"解决方案
// 只有当 "空格空格" 的时候, 才做删除, 当前读取的空格, 其他时候跳过
func squashSpace2(b []byte) []byte {

	// 我在外面控制i++, 没有内部手动控制, 会产生一个问题: 只能判断两个连续的空格, 删除一个空格后, 后面还有空格会往左移, 外面的i++跳过了检查
	// for i := 0; i < len(b); i++ {
	// 	if b[i] == ' ' && b[i-1] == ' ' {
	// 		b = append(b[:i], b[i+1:]...)
	// 	}
	// }

	// 在里面控制i++, 做了左移复制之后, 就不做i++, 继续检查后面的是不是空格. 避免了在外面做i++, 做了左移复制后, 就跳过了当前的空格
	for i := 1; i < len(b); {
		if b[i] == ' ' && b[i-1] == ' ' {
			b = append(b[:i], b[i+1:]...)
		} else {
			i++
		}
	}

	return b
	// 最终也是要返回b的,如果不返回, 原字符串后面会有很多的最后一位Hello world! ||||||, 是因为此时的b已经是一个子切片, 有自己新的len和新的cap
	// 如果打印还用原来的b, 原来的b标头值中是老的len, 老的cap, 看到的结果是不一样的
	// 根因解释：切片的“底层数组”没有变，长度变了但容量没变
	// 我们先看下面这段代码：
	// b := []byte("Hello   world!")
	// b = squashSpace2(b)
	// fmt.Println(string(b))
	// 你处理后得到了一个新切片（子切片），例如 []byte("Hello world!")，长度是 12，但它的 底层数组容量可能还是原来的 16（原来有 3 个多余空格）。
	// 现在关键点来了：
	// 如果你 不返回处理后的子切片，直接打印原 b：
	// squashSpace2(b)     // 改变的是内部 b 的内容，但原 b 的 len 没变
	// fmt.Println(string(b)) // 依然是原切片的 len 和 cap
	// 比如：
	// 原始 b：长度 16，容量 16
	// 去重后你得到的子切片 b[:12]，长度变成了 12
	// 但你没有返回它，而是打印原始的 b，它长度还是 16，那么后面的 4 个字节（原本是 ' '，但现在可能是旧值或被挪过来的）仍然会被打印。
	// 	举个更明显的例子：
	// b := []byte("Hello   world!")
	// squashSpace2(b)
	// fmt.Printf("原始 b 长度: %d, 值: %q\n", len(b), b)
	// 输出：
	// 原始 b 长度: 16, 值: "Hello world!d!"
	// 这是因为 b 的长度没有变，仍然是 16 个字节，只是前面一部分内容被修改，后面的是旧数据或挪动过程中复制过来的脏数据。

}

// 思路: 我们不能用 b[i] == ' ' 判断空格了，因为有很多其他 Unicode 空白字符，比如 \t、\n、全角空格等
// 1. 遍历整个 []byte 切片
// (1) 因为它是 UTF-8 编码，要用 utf8.DecodeRune 解码每个字符成 rune
// (2) 一次可能是 1-4 个字节
// 2. 判断这个字符是否是 Unicode 空白字符
// (1) 使用标准库函数：unicode.IsSpace(rune)
// 3. 这个字符是空白字符, 且前一个字符不是空白字符, 就要
// (1)
// 3. 跳过连续的空白字符，只保留一个 ASCII 空格（' '）
// 4. 把非空格字符原地写回切片
// (1) 用 copy() 或 utf8.EncodeRune() 写回 UTF-8 字节

func squashUnicodeSpaces(b []byte) []byte {
	writeIdx := 0
	prevIsSpace := false

	for readIdx := 0; readIdx < len(b); {
		// 每次遍历一个rune文字字符, 拿到这个rune文字字符的值, 长度(多个字节)
		r, size := utf8.DecodeRune(b[readIdx:])

		// 判断这个rune文字字符是否空白字符
		if unicode.IsSpace(r) {
			if !prevIsSpace { // 第一次就遇到空白字符
				b[writeIdx] = ' '
				writeIdx++
				prevIsSpace = true
			}
		} else {
			n := utf8.EncodeRune(b[writeIdx:], r)
			writeIdx += n
			prevIsSpace = false
		}
		readIdx += size

	}
	return b[:writeIdx]
}

// 我的思路: readIdx, writeIdx
// Decoderune当前字节流, 拿到第一个rune文字字符, 和他的长度
// 1. 判断文字字符是空格, 且前一个也是空格, 不要, 仅仅readIdx增加, 继续拿下一个rune文字字符, preIsSpace依然是true
// 2. 判断文字字符是空格, 且前一个不是空格, 要, b[write]=' '把他变成单纯空格, write++, readIdx+size, preIsSpace=true
// 3. 判断文字字符不是空格, 不论前一个是啥, 都要, 用encoderune把这个文字字符编码到[write:]后, write+size, readIdx+size,  preIsSpace=false

func squashUnicodeSpaces1(b []byte) []byte {
	writeIdx := 0
	prevIsSpace := false

	// readIdx<len(b) 说明循环到字符串的头了
	for readIdx := 0; readIdx < len(b); {
		r, size := utf8.DecodeRune(b[readIdx:])
		// 当前字符是空白字符, 且前一个也是空白字符, 不要
		if unicode.IsSpace(r) && prevIsSpace == true {
			readIdx += size
		}
		// 当前字符是空白字符, 且前一个字符不是空白字符, 要
		if unicode.IsSpace(r) && prevIsSpace == false {
			b[writeIdx] = ' '
			readIdx += size
			writeIdx++
			prevIsSpace = true
		}
		// 当前字符不是空白字符, 不管前面字符是啥, 都要
		if !unicode.IsSpace(r) {
			// b[writeIdx] = byte(r) // 这么写不对, 此时r是rune文字字符, 我要编码成utf8放到b里面
			// 这么写不可以, unicode是4个字节编码, 转换成byte只会截取最低位的1个字节, 对于其他类型的文字字符会乱码

			// b[writeIdx:]语句的意思是: 不是追加的意思, 是把rune文字字符的unicode编码转成3个字节的utf8编码后, 放到从writeIdx索引开始的位置
			// encodeRune(b[:writeIdx],r), 如果这么写, 是把 rune文字字符的unicode编码转成3个字节的utf8编码后, 覆盖当前已经写好的部分
			runesize := utf8.EncodeRune(b[writeIdx:], r)
			readIdx += runesize
			writeIdx += runesize
			prevIsSpace = false
		}
	}
	return b[:writeIdx]
	// b[:writeIdx], 这个子切片的意思就是当前已经写好的部分
}
