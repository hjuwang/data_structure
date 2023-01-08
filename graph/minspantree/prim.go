package minspantree

import (
	"fmt"
	. "graph/datastruct/weightedgraph"
)

// 类型定义

// EdgeIndex 迫切需要学会反射和泛型
func EdgeIndex(edges []Edge, edge Edge) int {
	index := -1
	for key, value := range edges {
		if value == edge {
			index = key
		}
	}
	return index
}

// MiniSpanTreePrim 最小生成树算法
func MiniSpanTreePrim(G *Graph) {

	adjvex := make([]int, G.NumVex)  //保存相关顶点间边的权值点的下标
	lowcost := make([]int, G.NumVex) //保存相关顶点间边的权值
	lowcost[0] = 0                   //初始化第一个权值为0，即V0加入生成树，谁下标等于0谁加入生成树
	adjvex[0] = 0                    //初始化第一个顶点下标为0

	//初始化
	for i := 1; i < G.NumVex; i++ {
		lowcost[i] = G.Arc[0][i]
		adjvex[i] = 0 //初始化都为v0的下标
	}

	//整个大循环是构造最小生成树的过程（八天肯定能干完）
	for i := 1; i < G.NumVex; i++ {
		min, j, k := INF, 1, 0
		for j = 1; j < G.NumVex; j++ {
			//lowcost[j]=0已经是最小生成树中的顶点，不参与权值的查找
			if lowcost[j] != 0 && lowcost[j] < min {
				min = lowcost[j]
				//找到权值最小的顶点赋值给k
				k = j
			}
		}
		//每次输出一个最小生成树的边
		fmt.Printf("(%d,%d)", adjvex[k], k)
		lowcost[k] = 0
		for j = 1; j < G.NumVex; j++ {
			if lowcost[j] != 0 && G.Arc[k][j] < lowcost[j] {
				lowcost[j] = G.Arc[k][j]
				adjvex[j] = k
			}
		}
	}
}
