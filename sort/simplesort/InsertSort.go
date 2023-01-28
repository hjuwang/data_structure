package simplesort

// InsertSort 直接插入排序算法
func InsertSort(L *SqlList) {
	//在有序的序列中插入关键字
	for i := 2; i < L.length; i++ { //默认第一个数字是有序的序列
		if L.r[i] < L.r[i-1] { //需要将L.r[i]设置为哨兵准备插入
			L.r[0] = L.r[i] //设置哨兵
			var j int
			for j = i - 1; L.r[j] > L.r[0]; j-- {
				L.r[j+1] = L.r[j] //记录后移
			}
			L.r[j+1] = L.r[0] //哨兵插入
		}
	}
}
