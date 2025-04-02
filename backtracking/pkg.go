package main

type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v any) *TreeNode {

	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

// 切片构造二叉树,递归生成
/*
i 表示节点索引，root = 0
left = 2*i+1
right= 2*i+2

*/

func SliceToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i > len(arr)-1 || arr[i] == nil {
		return nil
	}

	root := NewTreeNode(arr[i])
	root.Left = SliceToTreeDFS(arr, 2*i+1)
	root.Right = SliceToTreeDFS(arr, 2*i+2)
	return root
}

func SliceToTree(arr []any) *TreeNode {
	return SliceToTreeDFS(arr, 0)
}

/*
将二叉树转换为切片
*/
func TreeToSliceDFS(root *TreeNode, i int, res *[]any) {
	// 根节点
	if root == nil {
		return
	}
	for i >= len(*res) {
		*res = append(*res, nil)
	}
	(*res)[i] = root.Val
	//处理左子树
	TreeToSliceDFS(root.Left, 2*i+1, res)
	//处理右子树
	TreeToSliceDFS(root.Right, 2*i+2, res)

}

func TreeToSlice(root *TreeNode) []any {
	var res []any
	TreeToSliceDFS(root, 0, &res)
	return res
}
