package balancebintree

func BalanceBinTree() {
	a := []int{3, 2, 1, 4, 5, 6, 7, 10, 9, 8}
	var T BiTree
	var taller bool
	T = nil
	for i := 0; i < len(a); i++ {
		InsertAVL(&T, a[i], &taller)
	}

}
