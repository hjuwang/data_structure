package simplesort

// SimpleSelectionSort 简单选择排序
func SimpleSelectionSort(L *SqlList) {
	for i := 0; i < L.length; i++ {
		min := i //假设最小值下标为i
		for j := i + 1; j < L.length; j++ {
			if L.r[min] > L.r[j] {
				min = j
			}
		}
		if min != i {
			Swap(L, i, min)
		}
	}
}

/*
对简单选择排序算法的解释：
通过n-i次比较
在n-i +1 个关键字中
找到最小的和i进行交换
*/
