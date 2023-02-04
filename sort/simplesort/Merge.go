package simplesort

// Merge 将有序的SR[i,m],SR[m+1,n],归并为有序的TR[i,n]中
func Merge(SR, TR []int, i, m, n int) {

	var k, j, l int
	for k, j = i, m+1; i <= m && j <= n; k++ {
		if SR[i] < SR[j] {
			TR[k] = SR[i]
			i++
		} else {
			TR[k] = SR[j]
			j++
		}
		if i <= m { //前半部分有剩余
			for l = 0; l < m-i; l++ {
				TR[k+l] = SR[i+l]
			}
		}
		if j <= n {
			for l = 0; l < n-j; l++ {
				TR[k+l] = SR[i+l]
			}
		}
	}
}
