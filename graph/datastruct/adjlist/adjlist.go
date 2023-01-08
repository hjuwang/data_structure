package adjlist

import "fmt"

// Edge 边定义（无权）
type Edge struct {
	Begin int
	End   int
}

// EdgeNode 边节点
type EdgeNode struct {
	Adjvex int
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
func ALInit(G *GraphAdjList, vex []int, edges []Edge) {
	G.NumVertexes = len(vex)
	G.NumEdges = len(edges)
	G.AdjList = make([]VertexNode, G.NumVertexes)
}

// ALCreate 创建邻接表（头差法）
func ALCreate(G *GraphAdjList, vex []int, edges []Edge) {
	//创建顶点集合
	for i := 0; i < G.NumVertexes; i++ {
		G.AdjList[i].In = 0
		G.AdjList[i].Data = vex[i]
		G.AdjList[i].FirstEdge = nil
	}

	//连接边
	for _, edge := range edges {
		beginVex := edge.Begin
		endVex := edge.End
		//创建边节点，并使用头插法插入节点
		edgeNode := new(EdgeNode)
		edgeNode.Adjvex = endVex
		edgeNode.Next = G.AdjList[beginVex].FirstEdge
		G.AdjList[beginVex].FirstEdge = edgeNode
		//添加顶点入度
		G.AdjList[endVex].In++
	}
}

// ALDeleteVex 删除邻接表中的顶点vex 位置是vexindex，注意删除元素会引起索引的变化
func ALDeleteVex(G *GraphAdjList, vexindex int) bool {

	var p *EdgeNode
	//判断位置非法
	if vexindex < 0 || vexindex > G.NumVertexes-1 {
		fmt.Println("位置非法")
		return false
	}
	//删除与顶点相关的入边
	for i := 0; i < len(G.AdjList); i++ {
		if G.AdjList[i].FirstEdge == nil {
			continue
		}
		//特殊处理第一个元素
		if G.AdjList[i].FirstEdge.Adjvex == vexindex {
			G.AdjList[i].FirstEdge = G.AdjList[i].FirstEdge.Next
			G.NumEdges--
			continue
		}
		//从链表中的第二个节点开始查找,找到并删除
		p = G.AdjList[i].FirstEdge
		for p.Next != nil {
			if p.Next.Adjvex == vexindex {
				p.Next = p.Next.Next
				G.NumEdges--
				break
			}
			p = p.Next
		}

	}
	//删除顶点的出边
	for e := G.AdjList[vexindex].FirstEdge; e != nil; e = e.Next {
		//自己的出边会影响其他顶点的in
		for i := 0; i < len(G.AdjList); i++ {
			if i == e.Adjvex {
				G.AdjList[i].In--
			}
		}
		G.NumEdges--
	}
	//删除顶点
	G.AdjList = append(G.AdjList[:vexindex], G.AdjList[vexindex+1:]...)
	//修改链表中的adjvex
	for i := 0; i < len(G.AdjList); i++ {
		for e := G.AdjList[i].FirstEdge; e != nil; e = e.Next {
			if e.Adjvex > vexindex {
				//向前移动一位
				e.Adjvex--
			}
		}
	}

	G.NumVertexes--
	return true

}

// VexIndex 根据内容获取位置
func VexIndex(G *GraphAdjList, data int) int {
	index := -1
	for i := 0; i < len(G.AdjList); i++ {
		if G.AdjList[i].Data == data {
			index = i
		}
	}
	return index
}

func AdjListView(G GraphAdjList) {
	fmt.Printf("Index\tIn\tData\tFirstEdge\t\n")
	index := 0
	for _, vex := range G.AdjList {
		fmt.Printf("[%d]\t%d\t%d\tFirstEdge-->", index, vex.In, vex.Data)
		for e := vex.FirstEdge; e != nil; e = e.Next {
			fmt.Printf("(Adjex:%d,Next-)->", e.Adjvex)
		}
		fmt.Printf("\n")
		index++
	}

}
