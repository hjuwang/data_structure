package shortestpath

import (
	"fmt"
	. "graph/datastruct/weightedgraph"
)

//迪杰斯特拉算法，计算顶点v到图中每个顶点的最短距离

var PathArc []int        //PathArc[k]=8 v到k的最短路径，k的前驱节点的下标是8
var ShortPathTable []int // ShortPathTable[k]=10 v到k点的最短距离是10
// Path 定义数组记录v到每个顶点的路径(顶点下标顺序)
var Path [][]int

// Dijkstra 算法实现
func Dijkstra(G *Graph, v0 int) {

	PathArc = make([]int, G.NumVex)
	ShortPathTable = make([]int, G.NumVex)
	final := make([]bool, G.NumVex) //final[k]=true 表示已经求的v0到k的最短路径
	k := -1                         //记录路径最短顶点的下标

	//初始化
	for v := 0; v < G.NumVex; v++ {
		final[v] = false                 //v0还没有找到任何一个最短的路径
		ShortPathTable[v] = G.Arc[v0][v] //与v0邻接点的权值
		PathArc[v] = -1                  //到每个点的路径还没有计算出，自然没有任何前驱，所以都初始化为-1
	}
	ShortPathTable[0] = 0 //v0到下标为0的顶点(v0)的距离为0
	final[0] = true       //v0到v0的最短路径已经求出

	//大循环计算vo到（v1-v8）中各个顶点的最短路径
	//第一次循环，计算v0到v1的最短路径，
	for v := 1; v < G.NumVex; v++ {
		//
		min := INF
		//寻找离v0最近的顶点
		for w := 0; w < G.NumVex; w++ {
			if !final[w] && ShortPathTable[w] < min {
				//记录离v0最近的顶点（的下标）
				k = w
				min = ShortPathTable[w]
			}
		}
		//找到以后在final中标记已经找到
		final[k] = true //目前来看v0 到vk的最短距离已经找到，

		//修改权值，如果v0-
		for w := 0; w < G.NumVex; w++ {
			if !final[w] && (min+G.Arc[k][w]) < ShortPathTable[w] {
				ShortPathTable[w] = min + G.Arc[k][w]
				//前驱设置为K
				PathArc[w] = k
			}
		}
	}
	PathPrint(4, PathArc)
}

func PathPrint(dest int, patharc []int) {
	var path []int
	for dest != -1 {
		path = append(path, dest)
		dest = PathArc[dest]
	}
	fmt.Println(path)
}
