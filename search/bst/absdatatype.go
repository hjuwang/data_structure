package bst

import "fmt"

// SearchBST 查找，p指向查找到的节点的位置,f指向双亲
func SearchBST(T BiTree, key int, f BiTree, p *BiTree) bool {
	if T == nil {
		*p = f
		return false
	} else if key == T.data {
		*p = T
		return true
	} else if key < T.data {
		return SearchBST(T.lchid, key, T, p)
	} else {
		return SearchBST(T.rchild, key, T, p)
	}
}

// InsertBST 插入元素,T是指向二叉链表的指针（指针的指针）
func InsertBST(T *BiTree, key int) bool {
	var p, s BiTree
	//如果查找不成功
	if !SearchBST(*T, key, nil, &p) {
		s = new(BiTNode)
		s.data = key
		s.lchid = nil
		s.rchild = nil
		if p == nil {
			*T = s
		} else if key < p.data {
			p.lchid = s
		} else {
			p.rchild = s
		}
		return true
	} else {
		return false
	}
}

// DeleteBST 删除二叉排序树中的节点
func DeleteBST(T *BiTree, key int) bool {
	if *T == nil {
		return false
	} else {
		if key == (*T).data {
			Delete(T)
		} else if key < (*T).data {
			DeleteBST((*BiTree)(&(*T).lchid), key)
		} else {
			DeleteBST((*BiTree)(&(*T).rchild), key)
		}
	}

	return true
}

// Delete 删除树中的节点p,并重接左右子树
func Delete(p *BiTree) bool {
	//定义两个可以指向树节点的指针变量
	var q, s BiTree
	if (*p).lchid == nil {
		//用左子树替换自己
		(*p) = (*p).lchid
	} else if (*p).rchild == nil {
		(*p) = (*p).rchild
	} else {
		//左右子树都不为空，找到前驱节点替换，并删除前驱节点
		q = *p
		s = (*p).lchid
		//在左子树中寻找最右侧的节点
		for s.rchild != nil {
			q = s
			s = s.rchild
		}
		//循环完以后，s指向最右侧节点（p的前驱节点），p指向s的双亲啊
		(*p).data = s.data
		if q != *p {
			q.rchild = s.lchid

		} else {
			//直接重接左子树
			q.lchid = s.lchid
		}

	}
	return true
}

// 中序遍历二叉树
func MiddleOrder(T BiTree) {
	if T != nil {
		MiddleOrder(T.lchid)
		fmt.Printf("%d,", T.data)
		MiddleOrder(T.rchild)
	}
}
