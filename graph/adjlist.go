package main

import (
	"bytes"
	"fmt"
)

// EdgeNode 使用邻接表实现图
// 边表节点
type EdgeNode struct {
	adjVex int       //邻接顶点的序号(数组下标)
	next   *EdgeNode //指向下一个边节点
}

// VertexNode 顶点
type VertexNode struct {
	data      byte
	firstEdge *EdgeNode
}

// GraphAdjList 邻接表数据结构
type GraphAdjList struct {
	adjList         []VertexNode //顶点集合
	numVex, numEdge int          //顶点数目和边数目

}

// GraphAdjListInit 初始化
func GraphAdjListInit(gl *GraphAdjList, vexListLen int) {
	gl.adjList = make([]VertexNode, vexListLen)
	gl.numVex = 0
	gl.numEdge = 0
}

// GraphAdjListCreat 创建邻接表
func GraphAdjListCreat(gl *GraphAdjList, vexes []byte, edges []edge) {
	gl.numVex = len(vexes)
	gl.numEdge = len(edges)
	//生成顶点表集合
	for i := 0; i < len(vexes); i++ {
		gl.adjList[i].data = vexes[i]
		gl.adjList[i].firstEdge = nil
	}
	//连接边表,使用头插法
	for _, edg := range edges {
		//获取边（vexi,vesj）在顶点数组中的下标作为序号
		i := bytes.IndexByte(vexes, edg.vex1)
		j := bytes.IndexByte(vexes, edg.vex2)
		//连接边节点，使用头插法
		edgi := new(EdgeNode)
		edgi.adjVex = j
		edgi.next = gl.adjList[i].firstEdge
		gl.adjList[i].firstEdge = edgi

		edgj := new(EdgeNode)
		edgj.adjVex = i
		edgj.next = gl.adjList[j].firstEdge
		gl.adjList[j].firstEdge = edgj

	}
}

// DFSGL 邻接表的深度优先遍历递归算法
func DFSGL(gl *GraphAdjList, i int) {
	var p *EdgeNode
	visited[i] = true
	fmt.Printf("%c,", gl.adjList[i].data)
	p = gl.adjList[i].firstEdge
	for p != nil {
		//如果没有被访问
		if !visited[p.adjVex] {
			DFSGL(gl, p.adjVex)
		}
		p = p.next
	}
}

// DFSGLTraverses 邻接表深度优先遍历
func DFSGLTraverses(gl *GraphAdjList) {
	visited = make([]bool, gl.numVex)
	for i := 0; i < gl.numVex; i++ {
		visited[i] = false
	}

	for i := 0; i < gl.numVex; i++ {
		if !visited[i] {
			DFSGL(gl, i)
		}
	}
}

// BFSGL 广度优先遍历邻接表
func BFSGL(gl *GraphAdjList) {
	//初始化
	queue := make([]int, 0, 15)
	visited = make([]bool, gl.numVex)
	for i := 0; i < gl.numVex; i++ {
		visited[i] = false
	}
	//访问顶点
	for i := 0; i < gl.numVex; i++ {
		if !visited[i] {
			//访问入队
			visited[i] = true
			fmt.Printf("%c,", gl.adjList[i].data)
			queue = append(queue, i)
			//访问邻接点
			for len(queue) != 0 {
				i = queue[0]
				queue = queue[1:]

				p := gl.adjList[i].firstEdge
				for p != nil {
					if !visited[p.adjVex] {
						visited[p.adjVex] = true
						fmt.Printf("%c,", gl.adjList[p.adjVex].data)
						queue = append(queue, p.adjVex)
					}
					p = p.next
				}
			}

		}
	}
}
