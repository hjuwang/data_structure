package main

func SelectSort(array []int) []int {

	for i := 0; i < len(array); i++ {

		// 记录最小元素的索引
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

/* 时间复杂度时等差数列的前 n 项 和

1+2+3+....n  , n^2 级别的时间复杂度
*/
