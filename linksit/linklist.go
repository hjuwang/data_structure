package main

import "fmt"

//定义节点类型

type LNode struct {
	data ElemType
	next *LNode
}

// InitLinkList 初始化单链表，创建头节点，data存储元素的个数，返回头指针
func InitLinkList() *LNode {
	head := new(LNode)
	head.data = 0
	return head
}

// LinkListEmpty 单链表是否为空
func LinkListEmpty(l *LNode) bool {
	if l.next == nil && l.data == 0 {
		return true
	}
	return false
}

// LinkListInsert 在单链表中插入元素,位置i(从一开始),头指针指向头节点
func LinkListInsert(l *LNode, i int, e ElemType) Status {
	//位置非法
	if i < 1 || i > int(l.data) {
		fmt.Print("位置非法")
		return Error
	}
	//p初始化为头指针，指向头节点，j表示i位置的前驱位置
	p := l
	j := 1
	//p指向位置j
	for p.next != nil && j < i {
		p = p.next
		j += 1
	}
	//创建节点
	s := new(LNode)
	s.data = e
	s.next = p.next
	p.next = s
	l.data += 1
	return OK
}

// LinkListGetElem 获取链表中第i个位置的元素,放到e中
func LinkListGetElem(lhead *LNode, i int, e *ElemType) Status {
	if i < 1 || i > int(lhead.data) {
		return Error
	}
	//扫描指针和位置下标
	p := lhead
	j := 1
	for p.next != nil && j < i {
		p = p.next
		j += 1
	}
	*e = p.next.data
	return OK

}

// LinkListLocateElem 返回给定元素的位置,返回0表示失败
func LinkListLocateElem(lhead *LNode, e ElemType) int {
	index := 0
	p := lhead
	j := 1
	for p.next != nil {
		p = p.next
		if p.data == e {
			index = j
			break
		}
		j += 1
	}

	return index
}

// LinkListDelete 删除元素,位置i,返回删除的数据
func LinkListDelete(lhead *LNode, i int, e *ElemType) Status {
	if i < 1 || i > int(lhead.data) {
		fmt.Print("位置非法")
		return Error
	}
	p := lhead
	j := 1
	//循环使p指向位置i-1的节点,循环i-1次
	for p.next != nil && j < i {
		p = p.next
		j += 1
	}
	//执行删除操作
	*e = p.next.data
	q := p.next
	p.next = q.next
	lhead.data -= 1

	return OK

}

// LinklistLength 返回单链表的长度
func LinklistLength(lhead *LNode) int {

	length := 0
	p := lhead

	//有一个算一个
	for p.next != nil {
		length += 1
		p = p.next
	}

	return length

}
