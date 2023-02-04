package simplesort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {

	MergeSortTest()

}

// 直接插入排序
func InsetSortTest() {
	demo := []int{5, 3, 4, 6, 2}
	var L SqlList
	InitSqlList(demo, &L)
	//直接插入排序
	InsertSort(&L)
	fmt.Println(L.r[1:])

}

// 希尔排序
func ShellSortTest() {

	demo := []int{9, 1, 5, 8, 3, 7, 4, 6, 2}
	var L SqlList
	InitSqlList(demo, &L)
	fmt.Printf("希尔排序前：\n")
	fmt.Println(L.r[1:])
	ShellSort(&L)
	fmt.Printf("希尔排序后：\n")
	fmt.Println(L.r[1:])

}

// 堆排序
func HeapSortTest() {
	demo := []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	var L SqlList
	InitSqlList(demo, &L)
	HeapSort(&L)
	fmt.Printf("堆排序后：\n")
	fmt.Println(L.r[1:])

}

// 非递归归并排序
func MergeSortTest() {
	demo := []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	var L SqlList
	InitSqlList(demo, &L)
	MergeSort(&L)
	fmt.Printf("归并序后：\n")
	fmt.Println(L.r[1:])

}
