package shortestpath

import (
	"fmt"
	. "graph/datastruct/weightedgraph"
)

// PathAr2d 存储路径
type PathAr2d [][]int

// ShortPathTable2d 存储权值
type ShortPathTable2d [][]int

var D ShortPathTable2d
var P PathAr2d

func Floyd(G *Graph) {
	//分配数组空间
	D = make([][]int, G.NumVex)
	P = make([][]int, G.NumVex)
	for i := 0; i < G.NumVex; i++ {
		D[i] = make([]int, G.NumVex)
		P[i] = make([]int, G.NumVex)
	}
	//初始化
	for v := 0; v < G.NumVex; v++ {
		for w := 0; w < G.NumVex; w++ {
			D[v][w] = G.Arc[v][w]
			P[v][w] = w
		}
	}
	//算法循环
	//每个下标轮流当中转路径
	for k := 0; k < G.NumVex; k++ {
		for v := 0; v < G.NumVex; v++ {
			for w := 0; w < G.NumVex; w++ {
				//如果经过k比之前的路径要小
				if D[v][k]+D[k][w] < D[v][w] {
					D[v][w] = D[v][k] + D[k][w]
					//路径设置为经过下标为k的顶点
					P[v][w] = P[v][k]
				}
			}
		}
	}

	//打印每个顶点的路径
	fmt.Println("打印每个顶点的路径")
	for v := 0; v < G.NumVex; v++ {
		for w := 0; w < G.NumVex; w++ {
			fmt.Printf("V%d<-->V%d 权值:%d", v, w, D[v][w])
			k := P[v][w]
			//起点
			fmt.Printf("路径：V%d->", v)
			for k != w {
				fmt.Printf("V%d->", k)
				k = P[k][w]
			}
			fmt.Printf("V%d\n", w)
		}
		fmt.Printf("\n")
	}
}
