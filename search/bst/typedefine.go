package bst

//数据结构定义

type BiTNode struct {
	data          int
	lchid, rchild *BiTNode
}

// BiTree 定义树变量(指向树的指针)
type BiTree *BiTNode
