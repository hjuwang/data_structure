package main

import "fmt"

func partition2(array []int, left, right int) int {

	fmt.Println(">>", array, left, right)
	//选取基准元素
	i, j := left, right
	base := array[left]
	// 直到 i 合 j 相遇
	for i < j {

		// 从右边往左边找到第一个小于base 的数字
		for i < j && array[j] >= base {
			j -= 1
		}
		// 从左往右找到第一个大于base 的数字
		for i < j && array[i] <= base {
			i += 1
		}
		//交换这两个数字的位置
		array[j], array[i] = array[i], array[j]
	}
	// 将基准数字放到边界
	array[i], array[left] = base, array[i]

	// 返回哨兵索引
	fmt.Println("<<", array, i)
	fmt.Println("----------")
	return i
}

func qkSort(array []int, left, right int) []int {
	if left >= right {
		return array
	}
	// 划分哨兵索引
	midbase := partition2(array, left, right)
	array = qkSort(array, left, midbase-1)
	array = qkSort(array, midbase+1, right)
	return array
}

func main() {
	array := []int{2, 4, 1, 0, 3, 5}
	sort := qkSort(array, 0, len(array)-1)
	fmt.Println("**", sort)
}
