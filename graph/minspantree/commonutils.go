package minspantree

type IntEdgeList []IntEdge

func (l IntEdgeList) Len() int {
	return len(l)
}

func (l IntEdgeList) Less(i, j int) bool {
	return l[i].weight < l[j].weight
}

func (l IntEdgeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
