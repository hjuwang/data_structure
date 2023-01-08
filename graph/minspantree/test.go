package minspantree

import . "graph/datastruct/weightedgraph"

// MiniTree 最小生成树
func MiniTree() {
	//顶点集合
	vexs := []VexType{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"}
	//边集合
	edges := []Edge{
		{"v0", "v1", 10},
		{"v0", "v5", 11},
		{"v1", "v2", 18},
		{"v1", "v6", 16},
		{"v1", "v8", 12},
		{"v2", "v3", 22},
		{"v2", "v8", 8},
		{"v3", "v8", 21},
		{"v3", "v6", 24},
		{"v3", "v7", 16},
		{"v3", "v4", 20},
		{"v4", "v5", 26},
		{"v4", "v7", 7},
		{"v5", "v6", 17},
		{"v6", "v7", 19},
	}

	var graph Graph
	WGInit(&graph, vexs)
	WGCreat(&graph, vexs, edges)
	//最小生成树kruskal算法
	MiniSpanTreekruskal(&graph)

}
