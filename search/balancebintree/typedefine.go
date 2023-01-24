package balancebintree

// BiTNode 平衡二叉树节点定义
type BiTNode struct {
	data           int
	bf             int //平衡因子
	lchild, rchild *BiTNode
}

type BiTree *BiTNode
