package main

import (
	"container/list"
	"fmt"
)

func main() {

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

/*
分治的经典问题，汉诺塔问题,
使用[]int 类型模拟从小到大的圆盘
*/

// 移动一个圆盘
func move(src, tar *list.List) {
	//从顶部拿出
	pan := src.Back()
	//押入另一个顶部
	tar.PushBack(pan.Value)
	src.Remove(pan)
}

// Hanoi i 表示src 中元素的个数
func Hanoi(i int, src, buf, tar *list.List) {

	if i == 1 {
		move(src, tar)
		return
	}
	Hanoi(i-1, src, tar, buf)
	move(src, tar)
	Hanoi(i-1, buf, src, tar)

}

func SolveHanoi(a, b, c *list.List) {
	n := a.Len()
	Hanoi(n, a, b, c)
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
