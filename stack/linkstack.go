package main

import "fmt"

// 栈的链式存储结构的定义
type stackNode struct {
	data ElemType   //存放的数据
	next *stackNode //指向下一个节点

}

// LinkStack 定义栈
type LinkStack struct {
	count int        //栈中元素的个数
	top   *stackNode //栈顶指针,指向栈中的第一个元素
}

// LSInit 初始化栈,(其实不用初始化，go语言会默认赋零值)
func LSInit(s *LinkStack) {
	s.count = 0
	s.top = nil
}

// LSPush 入栈操作
func LSPush(stack *LinkStack, e ElemType) {
	node := new(stackNode)
	node.data = e
	node.next = stack.top
	stack.top = node
	stack.count += 1
}

// LSPop 出栈操作,数据存储在e中
func LSPop(stack *LinkStack, e *ElemType) bool {
	if stack.top == nil {
		fmt.Println("栈空")
		return false
	}
	*e = stack.top.data
	//指针后移动
	stack.top = stack.top.next
	stack.count -= 1
	return true
}

// LSGetTop 返回栈顶元素
func LSGetTop(stack *LinkStack, e *ElemType) bool {
	if stack.top == nil {
		fmt.Println("栈空")
		return false
	}
	*e = stack.top.data
	return true
}
