package btree

import "testing"

func TestBtree(t *testing.T) {

	r := []int{22, 16, 41, 58, 8, 11, 12, 16, 17, 22, 23, 31, 41, 52, 58, 59, 61}

	var T BTree
	var s Result
	for i := 0; i < len(r); i++ {
		//查找位置
		s = SearchBtree(T, r[i])
		if s.tag == 0 {
			InsertBtree(&T, r[i], s.pt, s.i)
		}
	}
}
