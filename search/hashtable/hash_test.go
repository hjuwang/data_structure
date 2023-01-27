package hashtable

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	var data = []int{12, 67, 56, 16, 25, 37, 22, 29, 15, 47, 48, 34}
	var hashTable HashTable
	InitHashTable(&hashTable)
	//插入数据
	for _, value := range data {
		InsertHash(&hashTable, value)
	}

	//查找
	for _, value := range data {
		var addr int
		fmt.Printf("查找:%d,Succcess:%d\n", value, SearchHash(&hashTable, value, &addr))
	}
}
