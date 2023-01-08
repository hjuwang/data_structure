package tpsort

import (
	. "graph/datastruct/weightadjlist"
)

// TopLogicalSort 关键路径拓扑排序

func TopLogicalSort(G GraphAdjList) bool {
	count := 0
	for i := 0; i < G.NumVertexes; i++ {
		if G.AdjList[i].In == 0 {
			Push(&stack1, i)
		}
	}
	//初始化每件事件的开始事件都为0
	etv = make([]int, G.NumVertexes)
	for len(stack1) != 0 {
		gettop := Pop(&stack1)
		Push(&stack2, gettop)
		count++
		//并删除以这个节点为begin的边
		for e := G.AdjList[gettop].FirstEdge; e != nil; e = e.Next {
			adjvex := e.Adjvex
			G.AdjList[adjvex].In--
			if G.AdjList[e.Adjvex].In == 0 {
				Push(&stack1, e.Adjvex)
			}
			//adjvex 作为AdjList[gettop]的后续事件，所以etv[adjex]>=etv[gettop]+e.weight
			if etv[adjvex] < etv[gettop]+e.Weight {
				etv[adjvex] = etv[gettop] + e.Weight
			}
		}
	}
	return count == G.NumVertexes
}

// Push 入栈
func Push(stack *[]int, data int) {
	*stack = append(*stack, data)
}

// Pop 出栈
func Pop(stack *[]int) int {
	data := (*stack)[len(*stack)-1]
	//切掉最后一个元素
	*stack = (*stack)[:len(*stack)-1]
	return data
}
