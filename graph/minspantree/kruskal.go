package minspantree

import (
	"fmt"
	. "graph/datastruct/weightedgraph"
	"sort"
)

type IntEdge struct {
	begin, end int
	weight     int
}

func MiniSpanTreekruskal(G *Graph) {

	var IntEdges IntEdgeList
	IntEdges = make([]IntEdge, 0)
	parent := make([]int, G.NumVex)

	//创建边集并排序
	CreatEdges(G, &IntEdges)

	for i := 0; i < G.NumVex; i++ {
		//自己是自己的前驱
		parent[i] = i
	}
	//循环遍历每一条边
	for i := 0; i < G.NumEdge; i++ {
		n := FindRoot(&parent, IntEdges[i].begin)
		m := FindRoot(&parent, IntEdges[i].end)
		//如果边的两端返回的都是尾部且相等，则说明当前边会把生成树闭环
		if m != n {
			//将此边的节点放入下标为起点的parent中，表示此顶点加入生成树集合
			parent[n] = m
			fmt.Printf("(%d,%d) %d\n", IntEdges[i].begin, IntEdges[i].end, IntEdges[i].weight)
		}
	}

	//给IntEdges排序

}

// FindRoot 查找最小生成树的连线的尾部顶点下标
func FindRoot(parent *[]int, f int) int {

	for (*parent)[f] != f {
		f = (*parent)[f]
	}
	return f
}

func CreatEdges(G *Graph, intEdges *IntEdgeList) {

	//给IntEdges赋值
	for i := 0; i < G.NumVex; i++ {
		for j := i; j < G.NumVex; j++ {
			if G.Arc[i][j] != 0 && G.Arc[i][j] != INF {
				edg := IntEdge{begin: i, end: j, weight: G.Arc[i][j]}
				*intEdges = append(*intEdges, edg)
			}
		}
	}
	sort.Sort(*intEdges)
}
