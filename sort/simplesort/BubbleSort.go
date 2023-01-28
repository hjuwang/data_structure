package simplesort

// BubbleSort0 冒泡排序(初级版)
func BubbleSort0(L *SqlList) {
	for i := 0; i < L.length; i++ {
		for j := i + 1; j < L.length; j++ {
			if L.r[i] > L.r[j] {
				Swap(L, i, j)
			}
		}
	}
}

// BubbleSort 真正的冒泡排序算法
func BubbleSort(L *SqlList) {
	for i := 0; i < L.length; i++ {
		for j := L.length - 1; j > i; j-- {
			if L.r[j] > L.r[j+1] {
				Swap(L, i, j) //交换
			}
		}
	}
}

// BubbleSort2 改进的冒泡排序算法
func BubbleSort2(L *SqlList) {
	flag := true
	for i := 0; i < L.length && flag; i++ {
		flag = false
		for j := L.length - 1; j > i; j-- {
			Swap(L, i, j)
			flag = true
		}

	}
}
