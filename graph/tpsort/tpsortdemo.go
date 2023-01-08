package tpsort

import (
	"fmt"
	. "graph/datastruct/adjlist"
)

func TPSortDemo(G *GraphAdjList) bool {
	//记录numVex:
	numVex := G.NumVertexes
	//删除直到途中不再有入度为0的点
	count := 0
	existence := false
	//不知道循环多少次
	for true {
		existence = false
		for i := 0; i < len(G.AdjList); i++ {
			//找到以后输出并删除
			if G.AdjList[i].In == 0 {
				fmt.Printf("V%d,", G.AdjList[i].Data)
				ALDeleteVex(G, i)
				count++
				existence = true
			}
		}
		if !existence {
			break
		}
	}
	fmt.Printf("count:%d", count)
	if count < numVex {
		return false
	}
	return true
}
