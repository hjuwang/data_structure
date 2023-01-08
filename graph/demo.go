package main

import (
	"fmt"
)

// 类型定义

func Demo() {

	var a []int
	fmt.Println(len(a))

}

// AdjList 邻接表
func AdjList() {

	var adjList GraphAdjList

	vexes := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	//边集合
	edges := []edge{{'A', 'B'}, {'A', 'F'},
		{'B', 'C'}, {'B', 'I'}, {'B', 'G'},
		{'C', 'I'}, {'C', 'D'},
		{'D', 'I'}, {'D', 'G'}, {'D', 'H'}, {'D', 'E'},
		{'E', 'H'}, {'E', 'F'},
		{'F', 'G'},
		{'G', 'H'},
	}

	GraphAdjListInit(&adjList, len(vexes))
	GraphAdjListCreat(&adjList, vexes, edges)
	fmt.Println(adjList.numVex, adjList.numEdge)
	fmt.Println("邻接表-深度遍历")
	DFSGLTraverses(&adjList)
	fmt.Println("邻接表-广度遍历：")
	BFSGL(&adjList)
}

// AdjMatrix 邻接矩阵操作
func AdjMatrix() {
	var graphAdjMatrix MyGraph
	//顶点集合
	vexes := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	//边集合
	edges := []edge{{'A', 'B'}, {'A', 'F'},
		{'B', 'C'}, {'B', 'I'}, {'B', 'G'},
		{'C', 'I'}, {'C', 'D'},
		{'D', 'I'}, {'D', 'G'}, {'D', 'H'}, {'D', 'E'},
		{'E', 'H'}, {'E', 'F'},
		{'F', 'G'},
		{'G', 'H'},
	}
	AdjMatrixInit(&graphAdjMatrix, vexes)
	AdjMatrixCreat(&graphAdjMatrix, vexes, edges)
	fmt.Println("深度优先遍历：")
	DFSGMTraverses(&graphAdjMatrix)
	fmt.Println("广度优先遍历")
	BFSGM(&graphAdjMatrix)

}
