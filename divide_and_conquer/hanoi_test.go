package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestHanoi(t *testing.T) {
	// 列表尾部是柱子顶部
	A := list.New()
	for i := 5; i > 0; i-- {
		A.PushBack(i)
	}
	B := list.New()
	C := list.New()
	fmt.Println("初始状态下：")
	fmt.Print("A = ")
	PrintList(A)
	fmt.Print("B = ")
	PrintList(B)
	fmt.Print("C = ")
	PrintList(C)
	//
	SolveHanoi(A, B, C)

	fmt.Println("圆盘移动完成后：")
	fmt.Print("A = ")
	PrintList(A)
	fmt.Print("B = ")
	PrintList(B)
	fmt.Print("C = ")
	PrintList(C)
}

func PrintList(list *list.List) {
	if list.Len() == 0 {
		fmt.Print("[]\n")
		return
	}
	e := list.Front()
	// 强转为 string, 会影响效率
	fmt.Print("[")
	for e.Next() != nil {
		fmt.Print(e.Value, " ")
		e = e.Next()
	}
	fmt.Print(e.Value, "]\n")
}
