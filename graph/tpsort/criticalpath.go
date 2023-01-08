package tpsort

import (
	"fmt"
	. "graph/datastruct/weightadjlist"
)

// 保存各事件的最早开始时间和最晚开始时间
var etv, ltv []int

// 存储拓扑排序序列
var stack2 []int

// 存储入度为0的顶点
var stack1 []int

func DataEstablish() {

	//构造图的数据
	vex := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	edges := []Edge{
		{0, 1, 3}, {0, 2, 4},
		{1, 3, 5}, {1, 4, 6},
		{2, 3, 8}, {2, 5, 7},
		{3, 4, 3},
		{4, 6, 9}, {4, 7, 4},
		{5, 7, 6},
		{6, 9, 2},
		{7, 8, 5},
		{8, 9, 3},
	}
	var graph GraphAdjList
	var graphcopy GraphAdjList
	ALCreate(&graph, vex, edges)
	graphcopy.NumEdges = graph.NumEdges
	graphcopy.NumVertexes = graph.NumVertexes
	graphcopy.AdjList = make([]VertexNode, graph.NumVertexes)
	copy(graphcopy.AdjList, graph.AdjList)
	TopLogicalSort(graphcopy)
	fmt.Println("关键路径：")
	CriticalPath(graph)
}

//

func CriticalPath(G GraphAdjList) {
	//初始化各事件的最晚开始时间
	ltv = make([]int, G.NumVertexes)
	for i := 0; i < G.NumVertexes; i++ {
		ltv[i] = etv[G.NumVertexes-1]
	}
	//
	for len(stack2) != 0 {
		gettop := Pop(&stack2)
		//计算各最后事件的最晚开始事件
		for e := G.AdjList[gettop].FirstEdge; e != nil; e = e.Next {
			adjex := e.Adjvex
			if ltv[gettop] > ltv[adjex]-e.Weight {
				ltv[gettop] = ltv[adjex] - e.Weight
			}
		}
	}
	//打印活动(边)
	for i := 0; i < G.NumVertexes; i++ {

		for e := G.AdjList[i].FirstEdge; e != nil; e = e.Next {
			adjex := e.Adjvex
			ete := etv[i]                //开始时间
			lte := ltv[adjex] - e.Weight //结束时间
			if ete == lte {
				fmt.Printf("{<v%d,v%d>,weight:%d}\n", G.AdjList[i].Data, G.AdjList[adjex].Data, e.Weight)
			}
		}
	}

}
