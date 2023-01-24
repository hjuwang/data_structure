package binsorttree

func BinSortTree() {

	//生成二叉树 &tree是指向树的指针
	var tree BiTree
	a := []int{62, 88, 58, 47, 35, 73, 51, 99, 37, 93}
	//循环生成二叉排序树
	for _, value := range a {
		InsertBST(&tree, value)
	}

}
