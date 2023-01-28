package simplesort

// HeapSort 堆排序算法
func HeapSort(L *SqlList) {

	for i := L.length / 2; i > 0; i-- { //构建排序堆
		HeapAjust(L, i, L.length)
	}
	for i := L.length; i > 1; i-- {
		Swap(L, 1, i)
		HeapAjust(L, 1, i-1) //每次还剩下i-1个元素
	}

}

// HeapAjust 构建排序堆
func HeapAjust(L *SqlList, s int, m int) {
	temp := L.r[s] //设置根节点元素
	//循环遍历左孩子
	for j := s * 2; j <= m; j *= 2 {
		if j < m && L.r[j] < L.r[j+1] {
			j += 1
		}
		if temp > L.r[j] {
			break
		}
		//不大于就对调
		L.r[s] = L.r[j]
		s = j
	}
	L.r[s] = temp

}
