// 通过一次循环, 完成左旋. 通过环状替换算法实现
// 核心思想: 你把第一个元素搬到它旋转之后的位置, 把那个位置的元素在搬到它对应的位置, 直到回到起点为止
package main

func rotateLeft(s []int, k int) {
	// fmt.Printf("0 %p,%v\n", &s[1], s)
	// fmt.Printf("数字3的地址 %p\n", &s[2])
	// tmp := s[:n] // [1,2], tmp指向原始的底层数组1,2
	// fmt.Printf("1 %p,%v\n", &tmp[1], tmp)
	// s = s[n:] // [3 4 5 6], s也是原始切片. 子len是4, 子cap是4, 子cap是从新起点开始算起到底层数组的末尾
	// fmt.Printf("2 %p,%v len %d cap %d\n", s, s, len(s), cap(s))
	// s = append(s, tmp...)
	// fmt.Printf("3 %p,%v\n", s, s)

	// 原始索引：   0   1   2   3   4   5   6
	// 原始值：     1   2   3   4   5   6   7

	// 目标值：     4   5   6   7   1   2   3
	// 目标索引：   0   1   2   3   4   5   6

	// 原来的索引为 0 的值 1，要去目标索引 4
	// 原来的索引为 1 的值 2，要去目标索引 5
	// 原来的索引为 2 的值 3，要去目标索引 6
	// 原来的索引为 3 的值 4，要去目标索引 0
	// 原来的索引为 4 的值 5，要去目标索引 1
	// 原来的索引为 5 的值 6，要去目标索引 2
	// 原来的索引为 6 的值 7，要去目标索引 3

	// 规律: 每个位置的数字索引i, 新的位置是往后挪4个,
	// 原索引0 -> 新索引4
	// 原索引1 -> 新索引5
	// 原索引2 -> 新索引6
	// 原索引3 -> 新索引7  % 7 = 0
	// 原索引4 -> 新索引8  % 7 = 1
	// 原索引5 -> 新索引9  % 7 = 2
	// 原索引6 -> 新索引10 % 7 = 3

	// 新索引的公式: (i+(n-k))%n

	n := len(s)
	// 如果n=7, 左移8次, 8%7取余=1, 就相当于左移1次
	k = k % n
	// 统计完成多少个元素移动
	count := 0

	// startidx=0从索引0开始
	for startidx := 0; count < n; startidx++ {
		// currentidx是当前处理的位置, prev是当前这个位置的值, 这个值要被换到目标位置(先保存)

		currentidx := startidx
		origval := s[startidx]
		for {
			// 核心公式: 计算当前位置往左旋转后的目标位置, 超出len后取余, 放在前面的索引, nextidx是目标的索引位置
			nextidx := (currentidx + (n - k)) % n
			// 把prev放入next位置, 把next原本值保存起来做下一轮的prev
			s[nextidx], origval = origval, s[nextidx]
			// 更新current为新的位置, 计数+1. s[next]=挪过来的值, prev变成了原有位置的值要被挪走
			currentidx = nextidx
			count++

			// 循环到startidx说明循环完了
			if startidx == currentidx {
				break
			}
		}
	}
	// (n,k)最大公约数=1, 例如(7,2), 外层循环只需要跑一次
	// 步骤	current	next	                        操作	                   新值存入位置
	// 1	0	   (0 + 7 - 3) % 7 = 4	           s[4] = 1, origval s[4]=5	  s[4] = 1
	// 2	4	   (4 + 7 - 3) % 7 = 1	           s[1] = 5, origval s[1]=2	  s[1] = 5
	// 3	1	   (1 + 7 - 3) % 7 = 5	           s[5] = 2, origval s[5]=6	  s[5] = 2
	// 4	5	   (5 + 7 - 3) % 7 = 2	           s[2] = 6, origval s[2]=3	  s[2] = 6
	// 5	2	   (2 + 7 - 3) % 7 = 6	           s[6] = 3, origval s[6]=7	  s[6] = 3
	// 6	6	   (6 + 7 - 3) % 7 = 3	           s[3] = 7, origval s[3]=4	  s[3] = 7
	// 7	3	   (3 + 7 - 3) % 7 = 0	           s[0] = 4，回到起点

	// (n,k)最大公约数=2, 例如(6,2), 数组会分成2个不相交的环

	// 外层第一次循环：startidx = 0
	// 初始化：

	// currentidx = 0
	// origval = s[0] = 1
	// 接下来我们搬环（环1）：

	// nextidx = (0 + (6-2)) % 6 = 4
	// → 把 1 放到位置4，记住 s[4] = 5 要搬走
	// → 结果：[1 2 3 4 1 6], origval = 5
	// → currentidx = 4

	// nextidx = (4 + 4) % 6 = 2
	// → 把 5 放到位置2，记住 s[2] = 3 要搬走
	// → 结果：[1 2 5 4 1 6], origval = 3
	// → currentidx = 2

	// nextidx = (2 + 4) % 6 = 0
	// → 把 3 放到位置0，记住 s[0] = 1
	// → 结果：[3 2 5 4 1 6], origval = 1
	// → currentidx = 0（回到起点，结束）

	// 第一个循环完成，处理了索引：0 → 4 → 2 → 0，3个元素搬完！

	// 外层第二次循环：startidx = 1
	// 初始化：
	// currentidx = 1
	// origval = s[1] = 2
	// 继续搬环（环2）：

	// nextidx = (1 + 4) % 6 = 5
	// → 把 2 放到位置5，记住 s[5] = 6
	// → 结果：[3 2 5 4 1 2], origval = 6
	// → currentidx = 5

	// nextidx = (5 + 4) % 6 = 3
	// → 把 6 放到位置3，记住 s[3] = 4
	// → 结果：[3 2 5 6 1 2], origval = 4
	// → currentidx = 3

	// nextidx = (3 + 4) % 6 = 1
	// → 把 4 放到位置1，记住 s[1] = 2
	// → 结果：[3 4 5 6 1 2], origval = 2
	// → currentidx = 1（回到起点，结束）

	// 第二个循环完成，处理了索引：1 → 5 → 3 → 1，又搬了3个！

}

func rotateRight(s []int, k int) {
	n := len(s)
	// 避免出现传参k>n
	k = k % n
	count := 0
	for startIdx := 0; count < n; startIdx++ {
		currentIdx := startIdx
		// 老值给临时变量
		tmpVal := s[startIdx]
		for {
			// 新位置
			nextIdx := (currentIdx + k) % n
			// 新位置用老值(临时变量), 原来的值给临时变量
			s[nextIdx], tmpVal = tmpVal, s[nextIdx]
			currentIdx = nextIdx
			count++
			if currentIdx == startIdx {
				break
			}
		}
	}

}

// 原始索引：   0   1   2   3   4   5   6
// 原始值：     1   2   3   4   5   6   7

// 目标值：     5   6   7   1   2   3   4
// 目标索引：   0   1   2   3   4   5   6

// 原始索引0    新位置索引3
// 原始索引1    新位置索引4
// 原始索引2    新位置索引5
// 原始索引3    新位置索引6
// 原始索引4    新位置索引 4+3%7 = 0
// 原始索引5    新位置索引 5+3%7 = 1
// 原始索引6    新位置索引 6+3%7 = 2

// 右旋的公式: (i+k)%n
