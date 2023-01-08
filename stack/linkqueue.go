package main

import "fmt"

// QueueNode 节点类型
type QueueNode struct {
	data ElemType
	next *QueueNode
}

// LinkQueue QueueNode 队列的链式存储数据类型
type LinkQueue struct {
	front, rear *QueueNode
}

// Init 链式队列初始化
func (queue *LinkQueue) Init() {
	//定义队头节点
	node := new(QueueNode)
	node.next = nil
	queue.front = node
	queue.rear = node

}

// EnLinkQueue 入队操作
func (queue *LinkQueue) EnLinkQueue(e ElemType) {
	endNode := new(QueueNode)
	endNode.data = e
	//q.rear指向队尾节点
	queue.rear.next = endNode
	queue.rear = endNode
}

// DeLinkQueue 出队操作
func (queue *LinkQueue) DeLinkQueue(e *ElemType) bool {
	//判断队空
	if queue.rear == queue.front {
		fmt.Println("队空")
		return false
	}
	//指向队列中的第一个元素
	startNode := queue.front.next
	*e = startNode.data
	//头节点的next指向第二个节点
	queue.front.next = startNode.next
	//处理队尾指针
	if queue.rear == startNode {
		queue.rear = queue.front
	}
	return true

}

// QueueEmpty 判断队列是否是为空
func (queue *LinkQueue) QueueEmpty() bool {

	//其实这里是有小漏洞的
	switch true {
	case queue.front == queue.rear:
		return true
	case queue.front.next == nil:
		return true
	default:
		return false
	}

}

// QueueLength 返回队列元素的个数
func (queue *LinkQueue) QueueLength() int {
	count := 0
	q := queue.front.next
	for q != nil {
		count += 1
		q = q.next
	}
	return count
}
