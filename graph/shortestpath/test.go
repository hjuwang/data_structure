package shortestpath

import (
	. "graph/datastruct/weightedgraph"
)

// ShortestPath 最短路径算法
func ShortestPath() {

	vexs := []VexType{"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"}
	edges := []Edge{
		{"v0", "v1", 1},
		{"v0", "v2", 5},
		{"v1", "v2", 3},
		{"v1", "v3", 7},
		{"v1", "v4", 5},
		{"v2", "v4", 1},
		{"v2", "v5", 7},
		{"v3", "v4", 2},
		{"v4", "v5", 3},
		{"v3", "v6", 3},
		{"v4", "v6", 6},
		{"v4", "v7", 9},
		{"v5", "v7", 5},
		{"v6", "v7", 2},
		{"v6", "v8", 7},
		{"v7", "v8", 4},
	}
	//创建加权图
	var wg Graph
	WGInit(&wg, vexs)
	WGCreat(&wg, vexs, edges)
	Floyd(&wg)
}
