package main

import "fmt"

// MAXSIZE 定义数组的最大长度
const MAXSIZE = 100

type SqStack struct {
	data [MAXSIZE]ElemType //数组用于存放数据
	top  int               //栈顶指针 -1为空栈，正好与数组的下标对应
}

// InitSqStack 初始化栈，(更像是格式化栈)
func InitSqStack(stack *SqStack) {
	for i := 0; i < MAXSIZE; i++ {
		stack.data[i] = 0
	}
	stack.top = -1
}

// Push 入栈操作
func Push(stack *SqStack, e ElemType) bool {
	if stack.top == MAXSIZE-1 {
		fmt.Println("栈满")
		return false
	}
	//不满
	stack.top += 1
	stack.data[stack.top] = e
	return true
}

// Pop 出栈操作,将数据存在e中
func Pop(stack *SqStack, e *ElemType) bool {
	if stack.top == -1 {
		fmt.Println("栈空")
		return false
	}
	*e = stack.data[stack.top]
	//回收数组空间
	stack.data[stack.top] = 0
	stack.top -= 1
	return true
}

// GetTop 用e,返回栈顶元素
func GetTop(stack *SqStack, e *ElemType) bool {
	if stack.top == -1 {
		fmt.Println("栈空")
		return false
	}
	*e = stack.data[stack.top]
	return true
}
