package hashtable

// InitHashTable 初始化散列表
func InitHashTable(H *HashTable) bool {
	m = HASHSIZE
	H.count = m
	H.elem = make([]int, m)
	for i := 0; i < m; i++ {
		H.elem[i] = NULLKEY
	}
	return true
}

// Hash 散列函数，除留余数法
func Hash(key int) int {

	return key % m
}

// InsertHash 散列表插入关键字
func InsertHash(H *HashTable, key int) {
	addr := Hash(key)
	for H.elem[addr] != NULLKEY { //如果冲突
		addr = (addr + 1) % m //开放定址法寻址
	}
	H.elem[addr] = key
}

// SearchHash 散列表查找,找到的地址存储在addr中
func SearchHash(H *HashTable, key int, addr *int) int {
	*addr = Hash(key)
	for H.elem[*addr] != key { //如果冲突
		*addr = ((*addr) + 1) % m                           //开放探测寻址
		if H.elem[*addr] == NULLKEY || *addr == Hash(key) { //如果始终都没有找到
			return UNSUCCESS
		}

	}
	return SUCCESS
}
