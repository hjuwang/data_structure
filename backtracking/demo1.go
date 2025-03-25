package main

import "fmt"

type TreeNode struct {
	V int
	L *TreeNode
	R *TreeNode
}

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

func main() {
	var path []int
	var res [][]int

	preOrder(Tree, &path, &res)
	fmt.Println(res)
}

func preOrder(node *TreeNode, path *[]int, res *[][]int) {
	if node == nil {
		return
	}
	*path = append(*path, node.V)
	if node.V == 7 { // 说明找到了路径
		*res = append(*res, *path)
	}
	preOrder(node.L, path, res)
	preOrder(node.R, path, res)
	// 路径退回到之前
	*path = (*path)[:len(*path)-1]
}
