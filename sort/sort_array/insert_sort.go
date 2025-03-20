package main

/*
像排队一样，占领你的位置，后面的人往后走
*/
func InsertSort(array []int) []int {

	for i := 1; i < len(array); i++ {
		// 依次选取base 放到正确的位置
		base := array[i]

		// 从base的前一个位置开始往后挪一个位置,
		//注意在已经排序的部分比较大小,
		// 比base 大的应该放在base 的后面
		j := i - 1
		// todo 这里尤其注意，一定要注意先确保 j>=0
		for j >= 0 && array[j] > base {
			array[j+1] = array[j]
			j -= 1
		}
		// 挪的操作结束的时候 array[j]<=base,array[j] 没有被挪走，array[j+1]被挪走了
		// 所以吧 base 放在空闲的位置---array[j+1]
		array[j+1] = base

	}
	return array
}
