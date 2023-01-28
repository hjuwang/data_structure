package simplesort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {

	HeapSortTest()

}

// 直接插入排序
func InsetSortTest() {
	demo := []int{0, 5, 3, 4, 6, 2}
	var L SqlList
	InitSqlList(demo, &L)
	//直接插入排序
	InsertSort(&L)
	ViewSqlList(&L)

}

// 希尔排序
func ShellSortTest() {

	demo := []int{0, 9, 1, 5, 8, 3, 7, 4, 6, 2}
	var L SqlList
	InitSqlList(demo, &L)
	fmt.Printf("希尔排序前：\n")
	ViewSqlList(&L)
	ShellSort(&L)
	fmt.Printf("希尔排序后：\n")
	ViewSqlList(&L)
}

// 堆排序
func HeapSortTest() {
	demo := []int{0, 50, 10, 90, 30, 70, 40, 80, 60, 20}
	var L SqlList
	InitSqlList(demo, &L)
	HeapSort(&L)
	fmt.Printf("堆排序后：\n")
	ViewSqlList(&L)

}
