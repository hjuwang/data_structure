package main

import "fmt"

func main() {

	list := []int{12, 45, 67, 1, 3, 4, 78, 34, 6}
	sort := quickSort(list, 0, len(list)-1)
	fmt.Println(sort)
}
