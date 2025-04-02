package main

// 判断当前状态是否是解
func isSolution(state *[]*TreeNode) bool {
	//
	return len(*state) != 0 && (*state)[len(*state)-1].Val == 7

}

// 将状态记录到解
func recordSolution(state *[]*TreeNode, res *[][]*TreeNode) {
	*res = append(*res, *state)
}

// 判断当前节点是否合法
func isValid(choice *TreeNode) bool {
	return choice != nil
}

// 更新状态
func makeChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = append(*state, choice)
}

// 回退状态
func undoChoice(state *[]*TreeNode) {
	*state = (*state)[:len(*state)-1]
}

func backTrackIII(state, choices *[]*TreeNode, res *[][]*TreeNode) {
	// 检查是否为解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
	}
	//遍历所有选择
	for _, choice := range *choices {
		// 如果尝试选择成功
		if isValid(choice) {
			//记录尝试
			makeChoice(state, choice)
			//记录下一轮选择
			temp := []*TreeNode{choice.Left, choice.Right}
			backTrackIII(state, &temp, res)

			//回退选择
			undoChoice(state)

		}
	}
}
