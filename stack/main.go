package main

import "fmt"

func main() {
	//顺序栈
	//var stack SqStack
	//fmt.Printf("%p", &stack)
	//InitSqStack(&stack)
	////入栈
	//Push(&stack, 11)
	//Push(&stack, 12)
	//Push(&stack, 13)
	//var top ElemType
	//var popget ElemType
	//Pop(&stack, &popget)
	//fmt.Println(popget)
	//GetTop(&stack, &top)
	//fmt.Println(top)

	//链栈
	// {
	//	var ls LinkStack
	//	LSPush(&ls, 11)
	//	var popget ElemType
	//	LSPop(&ls, &popget)
	//	fmt.Println(popget)
	//	LSPop(&ls, &popget)
	//}

	var queue LinkQueue
	queue.Init()
	queue.EnLinkQueue(11)
	queue.EnLinkQueue(11)
	queue.EnLinkQueue(11)
	queue.EnLinkQueue(11)

	fmt.Println(queue.QueueEmpty())
	fmt.Println(queue.QueueLength())

}
