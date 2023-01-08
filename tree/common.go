package main

type hfNOde struct {
	data           int
	lchild, rchild *hfNOde
}

type hfNodeSlice []*hfNOde

// Len 根据节点的类型实现排序需要的方法
func (h hfNodeSlice) Len() int {
	return len(h)
}
func (h hfNodeSlice) Less(i, j int) bool {
	return h[i].data < h[j].data
}
func (h hfNodeSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func MaxInt(a, b int) int {
	max := a
	if a < b {
		max = b
	}
	return max
}
