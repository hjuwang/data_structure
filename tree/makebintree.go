package main

import (
	"fmt"
)

//定义二叉树结构

type ElemType string

// BiTNode 定义二叉树结构
type BiTNode struct {
	Data           ElemType
	LChild, RChild *BiTNode
}

func CreatBiTree() *BiTNode {

	H := &BiTNode{"H", nil, nil}
	G := &BiTNode{"G", nil, nil}
	D := &BiTNode{"D", G, H}
	I := &BiTNode{"I", nil, nil}
	E := &BiTNode{"E", nil, I}
	F := &BiTNode{"F", nil, nil}
	C := &BiTNode{"C", E, F}
	B := &BiTNode{"B", D, nil}
	A := &BiTNode{"A", B, C}
	return A
}

// 前序遍历函数

// PreOrder 前序便利
func PreOrder(tree *BiTNode) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Data)
	PreOrder(tree.LChild)
	PreOrder(tree.RChild)

}

// InOrder 中序遍历
func InOrder(tree *BiTNode) {
	if tree == nil {
		return
	}
	InOrder(tree.LChild)
	fmt.Print(tree.Data)
	InOrder(tree.RChild)

}

// PostOrder 后续遍历
func PostOrder(tree *BiTNode) {
	if tree == nil {
		return
	}
	PostOrder(tree.LChild)
	PostOrder(tree.RChild)
	fmt.Print(tree.Data)

}

// DeepOrder 层序遍历
func DeepOrder(root *BiTNode) {
	if root == nil {
		return
	}
	//打印当前曾的数据
	fmt.Print(root.Data)
	fmt.Print(root.LChild.Data)
	fmt.Print(root.RChild.Data)

}

// InitTree 构造空树
func InitTree(t *BiTNode) {
	t = nil
}

// TreeEmpty 判断空树
func TreeEmpty(t *BiTNode) bool {
	flag := false
	if t == nil {
		flag = true
	}
	return flag
}

// TreeDepth 返回树的深度
func TreeDepth(t *BiTNode) int {
	if t == nil {
		return 0
	}
	return MaxInt(TreeDepth(t.LChild), TreeDepth(t.RChild)) + 1

}
