package simplesort

// ShellSort 希尔排序
func ShellSort(L *SqlList) {

	var i, j int
	increment := L.length
	//先循环在判断
	for {
		increment = increment/3 + 1                //增量
		for i = increment + 1; i < L.length; i++ { //外部大循环
			if L.r[i] < L.r[i-increment] {
				L.r[0] = L.r[i]                                                   //暂存较小的数据
				for j = i - increment; j > 0 && L.r[0] < L.r[j]; j -= increment { //较大的数字向后移动,注意这里的判断条件
					L.r[j+increment] = L.r[j] //记录向后移动
				}
				L.r[j+increment] = L.r[0] //插入数据

			}

		}
		if increment <= 1 {
			break
		}
	}

}

//第一批突破n²时间复杂度的排序算法之一
