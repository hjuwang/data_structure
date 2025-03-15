package main

// 目的是找到分治的基本索引
func partition(list []int, left, right int) int {

	l, r := left, right

	//p =list[left]
	for l < r { // 知道左侧游标 和 右侧游标相遇为止
		for l < r && list[l] < list[left] {
			l += 1
		}
		// 找到左侧首个大于基准元素的元素
		for l < r && list[r] > list[left] {
			r -= 1
		}
		// 找到右侧首个小于基准元素的元素

		//交换
		list[l], list[r] = list[r], list[l]

	}
	//  此时交换基准元素和分支的边界
	list[l], list[left] = list[left], list[l]

	return l
}

func quickSort(list []int, left, right int) []int {
	if left >= right {
		return list
	}
	p := partition(list, left, right)
	list = quickSort(list, left, p-1)
	list = quickSort(list, p+1, right)
	return list
}
