package tpsort

import (
	"fmt"
	. "graph/datastruct/adjlist"
)

// TPSort 计数，如果环不存在，返回true ,否则返回false ,
func TPSort(G GraphAdjList) bool {
	count := 0
	gettop := 0
	//栈中只存储入度为0的元素
	var stack []int
	//入度为0的顶点入栈
	for i := 0; i < G.NumVertexes; i++ {
		if G.AdjList[i].In == 0 {
			stack = append(stack, i)
		}
	}
	//
	for len(stack) != 0 {
		//出栈(表示从图中删除元素)
		gettop = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		fmt.Printf("V%d,", G.AdjList[gettop].Data)
		for e := G.AdjList[gettop].FirstEdge; e != nil; e = e.Next {
			adjex := e.Adjvex
			G.AdjList[adjex].In--
			if G.AdjList[adjex].In == 0 {
				stack = append(stack, adjex)
			}
		}
	}
	//判断是否有回路
	if count < G.NumVertexes {
		return false
	} else {
		return true
	}
}

//将拓扑排序序列入栈
