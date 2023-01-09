package main

//对于有序表的查找算法

// 折半查找算法,查找元素的位置
func BinarySearch(a []int, key int) int {
	low := 1
	high := len(a)

	for low < high {
		//折半
		mid := (low + high) / 2
		if key < a[mid] {
			high = mid - 1
		} else if key > a[mid] {
			low = mid + 1
		} else {
			return mid
		}

	}
	return 0
}

//插值查找的算法

func InterPolationSearch(a []int, key int) int {
	low := 1
	high := len(a) - 1

	for low < high {
		//注意这个插值，插的很巧妙
		mid := low + (high-low)*(key-a[low])/a[high] - a[low]
		if key < a[mid] {
			high = mid - 1
		} else if key > a[mid] {
			low = mid + 1
		} else {
			return mid
		}

	}
	return 0

}
