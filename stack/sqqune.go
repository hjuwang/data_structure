package main

import "fmt"

const QUMAXSIZE = 10

// SqQueue 顺序循环队列定义
type SqQueue struct {
	data        [QUMAXSIZE]ElemType
	front, rear int
}

// SQQInit 初始化
func SQQInit(queue *SqQueue) {
	queue.rear = 0
	queue.front = 0
}

// SQQLength 返回循环队列中元素的个数
func SQQLength(queue *SqQueue) int {

	count := (queue.rear + QUMAXSIZE - queue.front) % QUMAXSIZE

	return count

}

// EnSQQ 入队
func EnSQQ(queue *SqQueue, e ElemType) bool {
	//判断队满
	if (queue.rear+1)%QUMAXSIZE == queue.front {
		fmt.Println("队满")
		return false
	}
	queue.data[queue.rear] = e
	queue.rear = (queue.rear + 1) % QUMAXSIZE
	return true
}

// DeSQQ 出队
func DeSQQ(queue *SqQueue, e *ElemType) bool {
	//判断队空
	if queue.rear == queue.front {
		fmt.Println("队空")
		return false
	}
	*e = queue.data[queue.front]
	queue.front = (queue.front + 1) % QUMAXSIZE
	return true
}
