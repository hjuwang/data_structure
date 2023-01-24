package btree

const m = 3 //b树的阶数

// BTNode B树节点类型
type BTNode struct {
	keynum int     //节点中关键字的个数
	parent *BTNode //指向双亲节点
	//关键字类型
	node [m + 1]struct { //0下标未用，所以是m+1
		key    int     //关键字向量
		ptr    *BTNode //子树指针向量
		recptr int     //记录指针向量，每个关键字可以索引一个记录
	}
}
type BTree *BTNode

// Result B树查找结果类型
type Result struct {
	pt  *BTNode //指向找到的节点
	i   int     //i属于[1,m] ，表示关键字在节点中的序号
	tag int     //1.查找成功，0.查找失败
}
