package main

import (
	"bytes"
	"fmt"
)

//使用邻接矩阵实现存储无向图

// MyGraph 定义数据结构
type MyGraph struct {
	vexes              []byte  //顶点数组
	edges              [][]int //边表
	numNodes, numEdges int
}

// AdjMatrixInit 邻接矩阵初始化
func AdjMatrixInit(g *MyGraph, vexes []byte) {
	g.vexes = make([]byte, len(vexes))
	//初始化边表
	g.edges = make([][]int, len(vexes))
	for i := 0; i < len(vexes); i++ {
		g.edges[i] = make([]int, len(vexes))
	}
	g.numNodes = 0
	g.numEdges = 0
}

// AdjMatrixCreat 初始化图
func AdjMatrixCreat(g *MyGraph, vexes []byte, edges []edge) {
	g.numNodes = len(vexes)
	g.numEdges = len(edges)
	//生成顶点数组
	for i := 0; i < g.numNodes; i++ {
		g.vexes[i] = vexes[i]
	}
	//根据编集生成邻接矩阵
	for _, edg := range edges {
		i := bytes.IndexByte(vexes, edg.vex1)
		j := bytes.IndexByte(vexes, edg.vex2)
		g.edges[i][j] = 1
		g.edges[j][i] = 1

	}

}

// 对图的深度优先遍历，打印内容

// DFSGM 访问图中的顶点数组中的第i个顶点,并对其邻接点进行递归访问（深度优先遍历算法）
func DFSGM(g *MyGraph, i int) {
	visited[i] = true
	fmt.Printf("%c,", g.vexes[i])
	for j := 0; j < g.numNodes; j++ { //循环访问临界节点 //g.edges[i][j]==1顶点j是顶点i的邻接点 //!visited[j] 顶点j没有被访问过
		if g.edges[i][j] == 1 && !visited[j] { //邻接，&& 未被访问
			DFSGM(g, j)
		}
	}
}

// DFSGMTraverses 深度优先遍历
func DFSGMTraverses(g *MyGraph) {
	//初始化访问标志
	visited = make([]bool, g.numNodes)
	for i := 0; i < g.numNodes; i++ {
		visited[i] = false
	}
	//循环访问顶点
	for i := 0; i < g.numNodes; i++ {
		//如果没有被访问过
		if !visited[i] {
			DFSGM(g, i)
		}
	}
}

// BFSGM 广度优先遍历算法，
func BFSGM(g *MyGraph) {
	queue := make([]int, 0, 15)
	visited = make([]bool, g.numNodes)
	for i := 0; i < g.numNodes; i++ {
		visited[i] = false
	}

	for i := 0; i < g.numNodes; i++ {
		if !visited[i] {
			//访问并出队
			visited[i] = true
			fmt.Printf("%c,", g.vexes[i])
			queue = append(queue, i)

			//访问邻接点（体现广度的思想）
			for len(queue) != 0 {
				//出队首元素,赋值给i
				i = queue[0]
				queue = queue[1:]
				//访问邻接点
				for j := 0; j < g.numNodes; j++ {
					if g.edges[i][j] == 1 && !visited[j] {
						visited[j] = true
						fmt.Printf("%c,", g.vexes[j])
						queue = append(queue, j)
					}
				}
			}

		}
	}

}

// GmLocateVex  返回图中顶点的位置，不存在返回-1
func GmLocateVex(g *MyGraph, vex byte) int {
	index := bytes.IndexByte(g.vexes, vex)
	return index
}

// GmGetVex  不存在返回0
func GmGetVex(g *MyGraph, v byte) byte {

	index := bytes.IndexByte(g.vexes, v)
	if index != -1 {
		return g.vexes[index]
	}
	return 0
}

// GmFirstAdjVex 返回顶点v的一个邻接顶点
func GmFirstAdjVex(g *MyGraph, v byte) byte {
	var adjvex byte
	locate := bytes.IndexByte(g.vexes, v)
	//存在顶点
	if locate != -1 {
		for i := 0; i < g.numNodes; i++ {
			if g.edges[locate][i] == 1 {
				adjvex = g.vexes[i]
			}
		}
	}
	return adjvex
}

// 返沪下一个邻接顶点
func GmNextAdjVex(g *MyGraph, v byte, w byte) byte {
	//如果任意一个不存在
	var nextAdjVex byte
	vLocate := GmLocateVex(g, v)
	wLocate := GmLocateVex(g, w)
	if vLocate == -1 || wLocate == -1 {
		return 0
	}
	//从下一个位置开始寻找
	for i := wLocate + 1; i < g.numNodes; i++ {
		if g.edges[vLocate][i] == 1 {
			nextAdjVex = g.vexes[i]
		}
	}
	return nextAdjVex
}

// GmInsertVex 在图g中添加新的顶点
func GmInsertVex(g *MyGraph, v byte) {
	//顶点集合要改变，边集合矩阵也要改变
	g.vexes = append(g.vexes, v)
	//增加一行0，每行增加一列
	vexrow := make([]int, g.numNodes)
	for i := 0; i < g.numNodes; i++ {
		vexrow[i] = 0
	}
	g.edges = append(g.edges, vexrow)
	for i := 0; i < g.numNodes; i++ {
		g.edges[i] = append(g.edges[i], 0)
	}
	g.numNodes += 1
}

// GmDeleteVex 删除图中的顶点v，及相关的弧
func GmDeleteVex(g *MyGraph, v byte) bool {
	vlocate := GmLocateVex(g, v)
	if vlocate == -1 {
		fmt.Println("顶点不存在")
		return false
	}
	//矩阵删除，先删除行，在删除列
	g.edges = append(g.edges[:vlocate], g.edges[vlocate+1:]...)
	//每行删除一列
	for i := 0; i < len(g.edges); i++ {
		g.edges[i] = append(g.edges[i][:vlocate], g.edges[i][vlocate+1:]...)
	}
	//删除顶点集合
	g.vexes = append(g.vexes[:vlocate], g.vexes[vlocate+1:]...)
	g.numNodes -= 1
	return true
}

// GmInsertArc 在图中增加弧<v,w>,若是无向图，添加<w,v>
func GmInsertArc(g *MyGraph, v, w byte) bool {
	vLocate := GmLocateVex(g, v)
	wLocate := GmLocateVex(g, w)
	if vLocate == -1 || wLocate == -1 {
		return false
	}
	g.edges[vLocate][wLocate] = 1
	g.edges[wLocate][vLocate] = 1
	return true
}

// GmDeleteArc 在图中删除弧<v ，w>
func GmDeleteArc(g *MyGraph, v, w byte) bool {
	vLocate := GmLocateVex(g, v)
	wLocate := GmLocateVex(g, w)
	if vLocate == -1 || wLocate == -1 {
		return false
	}
	g.edges[vLocate][wLocate] = 0
	g.edges[wLocate][vLocate] = 0
	return true
}
