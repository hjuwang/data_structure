package main

import (
	"sort"
)

//生成哈夫曼树，返回树根的节点指针

func CreateHafuman(data []int) *hfNOde {
	//定义一个节点指针切片,如此以来就不好排序了
	var dataNode hfNodeSlice
	for _, dt := range data {
		nodePtr := &hfNOde{dt, nil, nil}
		dataNode = append(dataNode, nodePtr)
	}
	//对类型进行排序
	for len(dataNode) > 1 {

		sort.Sort(dataNode)
		root := new(hfNOde)
		l := dataNode[0]
		r := dataNode[1]
		root.data = l.data + r.data
		root.lchild = l
		root.rchild = r
		//删除前两个节点并添加新节点
		dataNode = dataNode[2:]
		dataNode = append(dataNode, root)
	}
	return dataNode[0]
}

// 前序遍历哈夫曼树
//func PreHfman(root interface{}) {
//	var hf *hfNOde
//	tp := reflect.ValueOf(root).Type()
//	tp2 := reflect.ValueOf(root)
//	fmt.Println()
//
//}
