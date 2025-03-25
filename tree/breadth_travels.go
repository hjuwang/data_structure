package main

import "fmt"

type TreeNode struct {
	V int
	L *TreeNode
	R *TreeNode
}

// 构造一颗二叉树
var Tree = &TreeNode{
	V: 1,
	L: &TreeNode{
		V: 7,
		L: &TreeNode{
			V: 4,
			L: nil,
			R: nil,
		},
		R: &TreeNode{
			V: 5,
			L: nil,
			R: nil,
		},
	},
	R: &TreeNode{
		V: 3,
		L: &TreeNode{
			V: 6,
			L: nil,
			R: nil,
		},
		R: &TreeNode{
			V: 7,
			L: nil,
			R: nil,
		},
	},
}

// 广度优先需要借助队列
var quene []*TreeNode

func main() {

	breadThTravel(Tree)

	// 这是一个二叉树

}

// 广度优先遍历的主要操作对象是队列
func breadThTravel(node *TreeNode) {

	// 初始化队列
	quene = []*TreeNode{node}
	// 只要队列不为空
	for len(quene) > 0 {
		//出队
		out := quene[0]
		quene = quene[1:]
		fmt.Println(out.V)
		// 左子入队
		if out.L != nil {
			quene = append(quene, out.L)
		}
		// 右子入队
		if out.R != nil {
			quene = append(quene, out.R)
		}

	}

}

/*
改算法的时间复杂度为 n ,因为总共有n 个 元素
*/
