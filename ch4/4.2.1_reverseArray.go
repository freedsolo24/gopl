// 使用数组指针代替slice
package main

// arr是指针类型的数组
func reverseArray(arr *[6]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		// 可以写成arr[i]:    通过指针类取索引拿值
		// 可以写成(*arr)[i]: (*arr)指向main函数的array变量, 在用array变量取索引拿值. 注意要有小括号.
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}
