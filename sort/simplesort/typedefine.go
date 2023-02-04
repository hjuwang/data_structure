package simplesort

const MAXSIZE = 10000

// SqlList 常用的排序顺序表
type SqlList struct {
	r      []int
	length int //记录顺序表的长度
}

// Swap 常用的交换函数,交换位置i,j的数据
func Swap(L *SqlList, i, j int) {
	L.r[i], L.r[j] = L.r[j], L.r[i]
}

// InitSqlList 顺序表初始化
func InitSqlList(demo []int, L *SqlList) {
	L.length = len(demo)
	L.r = make([]int, L.length+1)
	for i := 1; i < L.length+1; i++ { //从第一个开始复制
		L.r[i] = demo[i-1]
	}
}
