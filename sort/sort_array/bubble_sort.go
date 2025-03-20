package main

/*
冒泡排序算法实现数组排序
*/

func BubbleSort(array []int) []int {

	// 一定要注意这个索引问题
	for i := len(array) - 1; i > 0; i-- {

		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}

	return array
}

/*
优化的冒泡排序算法
*/
func BubbleSort2(array []int) []int {

	// 一定要注意这个索引问题
	var haveSwitch bool
	for i := len(array) - 1; i > 0; i-- {

		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j+1], array[j] = array[j], array[j+1]
				haveSwitch = true
			}
		}
		if !haveSwitch {
			break // 如果某轮冒泡没有交换 ，说明数组已经完成了排序，可以直接退出
		}
	}

	return array
}
