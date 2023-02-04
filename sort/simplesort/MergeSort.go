package simplesort

func MergeSort(L *SqlList) {

	var TR = make([]int, L.length) //申请额外空间
	k := 1
	for k > L.length {
		MergePass(L.r, TR, k, L.length)
		k = k * 2
		MergePass(TR, L.r, k, L.length)
	}

}

func MergePass(SR []int, TR []int, s int, n int) {
	var i, j int
	i = 1
	for i <= n-2*s+1 {
		Merge(SR, TR, i, i+s-1, i+2*s-1)
		i = i + 2*s //每次跳过两个有序序列的长度
	}
	if i < n-s+1 {
		Merge(SR, TR, i, i+s-1, n)
	} else {
		for j = i; j <= n; j++ {
			TR[j] = SR[j]
		}
	}
}
