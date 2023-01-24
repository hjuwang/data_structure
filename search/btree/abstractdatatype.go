package btree

import "fmt"

// Search 在根节点中寻找关键字K,,使得p->node[i].key≤K＜p->node[i+1].key */
func Search(P BTree, K int) int {
	j := 0
	for i := 1; i <= P.keynum; i++ {
		if P.node[i].key <= K {
			j = i
		}
	}
	return j
}

/*
在B树中查找关键字K
查找结果Resutl{*ptr ,i ,tag}
	成功：ptr指向K所在的节点，i=K在节点中的序号，tag=1
	不成功：K应该插入在ptr指向的节点中，位置在i,i+1, tag=0，找到这个位置并返回
*/

func SearchBtree(T BTree, K int) Result {

	var p, q BTree
	p = T   //p指向待查节点
	q = nil //q指向p的双亲
	found := false
	i := 0
	r := Result{}
	for p != nil && !found { //如果不为空且没有找到
		i = Search(p, K)
		if i > 0 && p.node[i].key == K {
			found = true //找到
		} else {
			q = p //自己赋值给双亲
			p = p.node[i].ptr
		}
	}
	r.i = i
	if found {
		r.pt = p
		r.tag = 1
	} else { //如果每找到，返回要插入节点的信息
		r.pt = q
		r.tag = 0
	}
	return r
}

// Insert /*将r->key、r和ap分别插入到q->key[i+1]、q->recptr[i+1]和q->ptr[i+1]中  ,r是一个记录，用一个节点表示
// @q 被插入的树
// @i 插入记录的前驱
// @key  插入记录的关键字
// @ap 插入的记录节点的孩子指针
func Insert(q *BTree, i int, key int, ap BTree) {
	for j := (*q).keynum; j > i; j-- { //最后一次循环j=i+1，所以能够空出第i+1个位置
		(*q).node[j+1] = (*q).node[j]
	}
	(*q).node[i+1].key = key
	(*q).node[i+1].recptr = key
	(*q).node[i+1].ptr = ap //指向孩子节点
	(*q).keynum++
}

// Split 将节点区分割成2个节点，后一半移入ap中
func Split(q *BTree, ap *BTree) {
	var i, s int
	s = (m + 1) / 2              //s=2
	*ap = new(BTNode)            //生成新节点ap
	(*ap).node[0] = (*q).node[s] //新节点的第一个关键字节点
	for i = s + 1; i <= m; i++ {
		(*ap).node[i-s] = (*q).node[i]
		if (*ap).node[i-s].ptr != nil { //如果存在孩子节点，则孩子节点的双亲指针指向自己
			(*ap).node[i-s].ptr.parent = *ap
		}
	}
	(*ap).keynum = m - s
	(*q).keynum = s - 1
	(*ap).parent = (*q).parent //q 和 ap 是兄弟
}

// NewRoot /* 生成含信息(T,r,ap)的新的根结点&T，原T和ap为子树指针 */
func NewRoot(T *BTree, key int, ap BTree) {
	var p BTree
	p = new(BTNode)              //创建一个新的节点，作为根节点
	p.node[0].ptr = *T           //原来的T作为新节点的子树
	*T = p                       //*T还是指向根节点
	if (*T).node[0].ptr != nil { //如果有孩子，请相认
		(*T).node[0].ptr.parent = *T
	}
	(*T).node[0].key = 1
	(*T).parent = nil
	(*T).keynum = 1
	(*T).node[1].key = key
	(*T).node[1].recptr = key
	(*T).node[1].ptr = ap
	//如果ap存在则给ap认亲
	if (*T).node[1].ptr != nil {
		(*T).node[1].ptr.parent = *T
	}
}

// InsertBtree  在m阶B树T上结点*q的key[i]与key[i+1]之间插入关键字K的指针r。若引起 */
/*  结点过大,则沿双亲链进行必要的结点分裂调整,使T仍是m阶B树。 */
func InsertBtree(T *BTree, key int, q BTree, i int) {
	var ap BTree
	finished := false
	ap = nil
	var s, rx int
	rx = key
	for q != nil && !finished { //如果节点q存在，且尚未完成插入操作
		Insert(&q, i, rx, ap)
		if q.keynum < m {
			finished = true
		} else { //节点过大，分裂节点q
			s = (m + 1) / 2
			rx = q.node[s].recptr
			Split(&q, &ap)
			q = q.parent
			if q != nil {
				i = Search(q, key) //在双亲中查找key插入的位置
			}
		}
		if !finished {
			NewRoot(T, rx, ap)
		}

	}
}

func Print(c BTree, i int) {
	fmt.Printf("%d", c.node[i].key)
}
