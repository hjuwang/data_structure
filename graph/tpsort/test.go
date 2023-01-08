package tpsort

import (
	. "graph/datastruct/adjlist"
)

func TestTpSort() {

	vex := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	edges := []Edge{
		{0, 4}, {0, 5}, {0, 11},
		{1, 2}, {1, 4}, {1, 8},
		{2, 5}, {2, 6}, {2, 9},
		{3, 2}, {3, 13},
		{4, 7},
		{5, 8}, {5, 12},
		{6, 5},
		{8, 7},
		{9, 10}, {9, 11},
		{10, 13},
		{12, 9},
	}

	var gal GraphAdjList
	var galCopy GraphAdjList

	//初始化和创建
	ALInit(&gal, vex, edges)
	ALCreate(&gal, vex, edges)
	ALInit(&galCopy, vex, edges)

	copy(galCopy.AdjList, gal.AdjList)
	galCopy.NumEdges = gal.NumEdges
	galCopy.NumVertexes = gal.NumVertexes

	TPSort(galCopy)
	AdjListView(gal)

}

//add wei
