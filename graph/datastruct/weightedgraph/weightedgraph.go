package weightedgraph

const INF = 999

type VexType string

// 定义边
type Edge struct {
	//顶点
	Vex1, Vex2 VexType
	//权值
	Weight int
}

// Graph 定义网数据结构
type Graph struct {
	Vexes           []VexType
	Arc             [][]int
	NumVex, NumEdge int
}

// WGInit 初始化
func WGInit(G *Graph, vexs []VexType) {
	G.Vexes = make([]VexType, len(vexs))
	G.Arc = make([][]int, len(vexs))
	for i := 0; i < len(vexs); i++ {
		G.Arc[i] = make([]int, len(vexs))
	}
	for i := 0; i < len(vexs); i++ {
		for j := 0; j < len(vexs); j++ {
			if i == j {
				G.Arc[i][j] = 0
			} else {
				G.Arc[i][j] = INF
			}

		}
	}
	G.NumVex = 0
	G.NumEdge = 0
}

func WGCreat(G *Graph, vexs []VexType, edges []Edge) {

	G.NumEdge = len(edges)
	G.NumVex = len(vexs)
	//顶点集
	for i := 0; i < G.NumVex; i++ {
		G.Vexes[i] = vexs[i]
	}
	//边集合
	for _, edge := range edges {
		v1index := VexIndex(vexs, edge.Vex1)
		v2index := VexIndex(vexs, edge.Vex2)
		G.Arc[v1index][v2index] = edge.Weight
		G.Arc[v2index][v1index] = edge.Weight

	}
}

// VexIndex Index 返回数组下标
func VexIndex(strs []VexType, str VexType) int {
	index := -1
	for key, value := range strs {
		if value == str {
			index = key
		}
	}
	return index
}
