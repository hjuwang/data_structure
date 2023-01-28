package simplesort

import "fmt"

const MAXSIZE = 10000

// SqlList 常用的排序顺序表
type SqlList struct {
	r      [MAXSIZE + 1]int
	length int //记录顺序表的长度
}

// Swap 常用的交换函数,交换位置i,j的数据
func Swap(L *SqlList, i, j int) {
	L.r[i], L.r[j] = L.r[j], L.r[i]
}

// InitSqlList 顺序表初始化
func InitSqlList(demo []int, L *SqlList) {
	L.length = len(demo)
	for key, _ := range demo {
		L.r[key] = demo[key]
	}
}

func ViewSqlList(L *SqlList) {
	fmt.Printf("[")
	for i := 0; i < L.length; i++ {
		fmt.Printf("%d,", L.r[i])
	}
	fmt.Printf("]")

}
