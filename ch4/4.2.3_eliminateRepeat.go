// 删除切片中重复的字符
// 算法: 双指针算法, 变换思路, 不用删除, 用双指针做覆盖
package main

func chRepeat1(s []string) []string {

	if len(s) == 0 {
		return s
	}
	// 第一个元素要保留, 肯定从索引1开始比较
	// 删除切片中的一个元素:
	// s := []string{"a", "b", "c", "d", "e"}  如果要删除idx=2的"c"
	// s[:2] → 得到前两个元素：["a", "b"]
	// s[3:] → 得到第 4 个到最后的元素：["d", "e"]
	// append(s[:2], s[3:]...) → 把后面这部分拼接到前面，就跳过了 "c"
	// 语句的意思是: 在s[:2]后面追加s[3:]
	for i := 1; i < len(s); {
		if s[i] == s[i-1] {
			s = append(s[:i], s[i+1:]...)
		} else {
			// 如果不手动控制i++,当发现两个元素相等时，你删除了 s[i],  删除后，后面的元素整体向左移动,但 i++ 仍然执行，你就直接跳过了下一个元素
			// 如果有多个连续重复的，比如 "a", "a", "a"，你只能删一个，会漏掉重复项
			i++
		}
	}
	return s
}

// 换个角度想, 不删除, 只保留, 用一个变量writeIdx记录保留了哪些项, 遍历整个切片s, 每当发现一个不重复的新元素, 就放到s[writeIdx]位置, 最后只保留前writeIdx项
func chRepeat2(s []string) []string {
	if len(s) == 0 {
		return s
	}
	// readIdx: 读取指针, writeIdx: 写入指针, 表示当前去重后的切片末尾的位置
	writeIdx := 1 // 从第一个元素开始写
	for readIdx := 1; readIdx < len(s); readIdx++ {
		// 索引0的元素永远在
		// 相同不处理, 不同说明是新的内容, 写入writeIdx的位置
		// 不重复就写入写指针处，同时写指针后移
		// 重复就跳过（不写入，不移动写指针）
		if s[readIdx] != s[writeIdx-1] {
			s[writeIdx] = s[readIdx]
			writeIdx++
		}
	}
	return s[:writeIdx]
}

// s := []string{"a", "a", "b", "b", "c", "c", "d"}

// 初始:
// readIdx=1, writeIdx=1
// s[1]="a", s[0]="a" -> 重复，跳过

// readIdx=2, s[2]="b", s[0]="a" -> 不重复，s[1]="b", writeIdx=2

// readIdx=3, s[3]="b", s[1]="b" -> 重复，跳过

// readIdx=4, s[4]="c", s[1]="b" -> 不重复，s[2]="c", writeIdx=3

// readIdx=5, s[5]="c", s[2]="c" -> 重复，跳过

// readIdx=6, s[6]="d", s[2]="c" -> 不重复，s[3]="d", writeIdx=4

// 最终结果:
// s[:4] = {"a", "b", "c", "d"}

func chRepeat3(s []string) []string {
	if len(s) == 0 {
		return s
	}

	writeIdx := 1
	for readIdx := 1; readIdx < len(s); readIdx++ {
		if s[readIdx] == s[writeIdx-1] {
			continue
		} else {
			s[writeIdx] = s[readIdx]
			writeIdx++
		}
	}
	return s[:writeIdx]
}
