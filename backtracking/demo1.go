package main

var Tree = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

//func main() {
//	var path []*TreeNode
//	var res [][]*TreeNode
//
//	preOrderII(Tree, &res, &path)
//	for i := range res {
//		for _, node := range res[i] {
//			fmt.Println(node.Val)
//		}
//	}
//}

/* 前序遍历：例题二 */
func preOrderII(root *TreeNode, res *[][]*TreeNode, path *[]*TreeNode) {
	if root == nil || root.Val.(int) == 3 {
		return
	}
	// 尝试
	*path = append(*path, root)
	if root.Val.(int) == 7 {
		// 记录解
		*res = append(*res, append([]*TreeNode{}, *path...))
	}
	preOrderII(root.Left, res, path)
	preOrderII(root.Right, res, path)
	// 回退,能走到这里，说明已经运行到了递归的归
	*path = (*path)[:len(*path)-1]
}
