package main

import "fmt"

func main() {
	//
	tree := CreatBiTree()

	fmt.Print("前序遍历:")
	PreOrder(tree)
	fmt.Println()
	fmt.Print("中序遍历:")
	InOrder(tree)
	fmt.Println()
	fmt.Print("后续遍历:")
	PostOrder(tree)
	fmt.Println()
	fmt.Print("层序遍历:")
	DeepOrder(tree)
	deep := TreeDepth(tree)
	fmt.Print("深度：", deep)

	//data := []int{5, 10, 15, 30, 40}
	//var root *hfNOde
	//root = CreateHafuman(data)
	////前序遍历
	//PreHfman(root)
	//定义

}
