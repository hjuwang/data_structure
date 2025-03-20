package main

import "fmt"

func main() {

	list := []int{12, 45, 67, 1, 3, 4, 78, 34, 6}
	//sort := quickSort(list, 0, len(list)-1)
	//fmt.Println(sort)

	// 失败
	//sort := BubbleSort2(list)
	//fmt.Println(sort)

	//sort = SelectSort(list)
	//fmt.Println(sort)

	insertSort := InsertSort(list)
	fmt.Println(insertSort)
}
