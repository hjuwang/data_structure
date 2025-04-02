package main

import "fmt"

func main() {

	//初始化二叉树人
	array := []any{1, 7, 3, 4, 5, 6, 7}

	root := SliceToTree(array)

	state := make([]*TreeNode, 0)
	res := make([][]*TreeNode, 0)
	// 一开始就是从根开始选择
	choice := []*TreeNode{root}

	backTrackIII(&state, &choice, &res)

	for _, path := range res {
		for _, node := range path {
			fmt.Printf("%v ", node.Val)
		}

	}
}
