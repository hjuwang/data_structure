package main

// MergeSort 递归地对数组进行归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

// merge 合并两个有序数组
/*
left 和 right 都是有序的
*/
func merge(left, right []int) []int {

	i, j := 0, 0
	result := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i]) // 每次追加的一定是两个有序数组中最小的
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// 看看哪个数组中有剩余的没参与比较的数组
	if i == len(left) {
		result = append(result, right[j:]...)
	} else {
		result = append(result, left[i:]...)
	}

	return result
}
