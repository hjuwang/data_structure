package simplesort

//快速排序

func QuickSort(L *SqlList){
	Qsort(L,1,L.length)
}

// Qsort 快速排序算法
func Qsort(L *SqlList,low ,high int){
	var pivot int
	if low <high{
		pivot =Partition(L,low,high)
		Qsort(L,low,pivot-1)
		Qsort(L,pivot+1,high)

	}
}

// Partition 返回序列的枢轴(序列左边都比枢轴小，右边都比枢轴大)
func Partition(L *SqlList, low ,high int) int{

	pivotkey := L.r[low] //初始化枢轴变量
	for low <high{
		for low < high &&  L.r[high]>=pivotkey{ //
			high--
		}
		Swap(L,low,high)
		for low < high &&  L.r[low]<=pivotkey{
			low++
		}
		Swap(L,low,high)
	}
	return low
}
