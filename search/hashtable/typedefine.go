package hashtable

const (
	SUCCESS   = 1
	UNSUCCESS = 0
	HASHSIZE  = 12 //散列表长为数组的长度
	NULLKEY   = -32768
)

type HashTable struct {
	elem  []int //存储数据，动态分配数组
	count int   //当前数据元素的个数
}

var m = 0 //散列表表长，全局变量
