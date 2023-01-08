package weightadjlist

import "fmt"

// Edge 带权有向图数据结构和抽象数据类型
// Edge 边定义（无权）
type Edge struct {
	Begin  int
	End    int
	Weight int
}

// EdgeNode EdgeNode 边节点
type EdgeNode struct {
	Adjvex int
	Weight int
	Next   *EdgeNode
}

// VertexNode 顶点节点
type VertexNode struct {
	In        int
	Data      int
	FirstEdge *EdgeNode
}

// GraphAdjList 邻接表数据结构定义
type GraphAdjList struct {
	AdjList               []VertexNode
	NumVertexes, NumEdges int
}

// ALInit 初始化
func ALInit(G *GraphAdjList, vex []int) {
	G.NumVertexes = 0
	G.NumEdges = 0
	G.AdjList = make([]VertexNode, len(vex))
}

// ALCreate 创建邻接表（头差法）
func ALCreate(G *GraphAdjList, vex []int, edges []Edge) {
	//创建顶点集合
	for _, value := range vex {
		//G.AdjList[i].In = 0
		//G.AdjList[i].Data = vex[i]
		//G.AdjList[i].FirstEdge = nil
		G.AdjList = append(G.AdjList, VertexNode{0, value, nil})
		G.NumVertexes++
	}

	//连接边
	for _, edge := range edges {
		beginVex := edge.Begin
		endVex := edge.End
		weight := edge.Weight
		//创建边节点，并使用头插法插入节点
		edgeNode := new(EdgeNode)
		edgeNode.Adjvex = endVex
		edgeNode.Weight = weight
		edgeNode.Next = G.AdjList[beginVex].FirstEdge
		G.AdjList[beginVex].FirstEdge = edgeNode
		//添加顶点入度
		G.AdjList[endVex].In++
		G.NumEdges++
	}
}

// AdjListView 邻接表可视化
func AdjListView(G GraphAdjList) {
	fmt.Printf("Index\tIn\tData\tFirstEdge\t\n")
	index := 0
	for _, vex := range G.AdjList {
		fmt.Printf("[%d]\t%d\t%d\tFirstEdge-->", index, vex.In, vex.Data)
		for e := vex.FirstEdge; e != nil; e = e.Next {
			fmt.Printf("(Adjex:%d,Weight:%d,Next-)->", e.Adjvex, e.Weight)
		}
		fmt.Printf("\n")
		index++
	}

}
