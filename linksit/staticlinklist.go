package main

import "fmt"

// SSLMAXSIZE 定义静态链表数据结构
const SSLMAXSIZE = 1000

type StaticLinkList struct {
	data ElemType //数据元素内容
	cur  int      //存放后继节点在数组中的下标
}

// SSL 定义静态链表类型
type SSL [SSLMAXSIZE]StaticLinkList

// InitSSL  初始化一个静态链表
// 第一个元素和最后一个元素不存数据
// 最后—个元素的cur则存放第-个有数值的元素的下标
func InitSSL(space *SSL) Status {

	for i := 0; i < SSLMAXSIZE-1; i++ {
		space[i].cur = i + 1
	}
	space[SSLMAXSIZE-1].cur = 0 //目前静态链表为空，最后一个元素的下标为0
	return OK

}

// MallocSSL 返回备用链表的下标，表满返回0
func MallocSSL(space *SSL) int {
	//数组的第一个元素的cur存储备用链表的第一个元素的下标，表满的space[0].cur=0
	i := space[0].cur
	//如果表不满，第一个元素重新分配备用链表
	if space[0].cur != 0 {
		space[0].cur = space[i].cur
	}
	return i
}

// FreeSSL 将下标为k的节点回收到备用链表
func FreeSSL(space *SSL, k int) {
	//k下标的位置回归空闲
	space[k].cur = space[0].cur
	//自己也回归空闲
	space[0].cur = k
}

// SSLLength 返回静态链表的长度(注意并非数组的长度)
func SSLLength(L *SSL) int {
	j := 0
	//获取链表中第一个元素的下标（在数组中的最后以元素存储）
	i := L[SSLMAXSIZE-1].cur
	//链表中的末尾元素的cur=0
	for i != 0 {
		i = L[i].cur
		j += 1
	}
	return j
}

//在静态链表的第i个位置插入元素

func SSLInsert(L *SSL, i int, e ElemType) Status {
	//判断非法
	if i < 1 || i > SSLLength(L)+1 {
		fmt.Print("位置非法")
		return Error
	}
	//获取备用链表的位置，表满返回0
	j := MallocSSL(L)
	//k首先是最后一个元素的下标
	k := SSLMAXSIZE - 1
	//判断链表未满
	if j != 0 {
		L[j].data = e
		for l := 1; l < i; l++ {
			k = L[k].cur
		}
		//循环完k是第i-1个元素的下标,把i-1的cur赋值给L[j].cur
		L[j].cur = L[k].cur
		L[k].cur = j
		return OK
	}
	return Error
}

// SSLDelete 静态链表的删除操作,将删除的元素，存储在e中
func SSLDelete(L *SSL, i int, e *ElemType) Status {
	//判断非法
	if i < 1 || i > SSLLength(L) {
		fmt.Print("位置非法")
		return Error
	}
	//判断表非空
	k := SSLMAXSIZE - 1
	//找到第i-1个元素的数组下标
	for l := 1; l < i; l++ {
		k = L[k].cur
	}
	//j为要删除元素的下标
	j := L[k].cur
	*e = L[j].data
	L[k].cur = L[j].cur
	FreeSSL(L, i)
	return OK
}

// SSLGetElem 获取链表中第i个位置的元素
func SSLGetElem(L *SSL, i int, desc *ElemType) Status {
	//判单位置非法
	if i < 1 || i > SSLLength(L) {
		fmt.Print("位置非法")
		return Error
	}
	k := SSLMAXSIZE - 1
	//循环i次k指向第i个元素
	for l := 1; l <= i; l++ {
		k = L[k].cur
	}
	*desc = L[k].data
	return OK
}

// SSLLocateElem 获取给定e的元素在链表中的位置,不存在返回0
func SSLLocateElem(L *SSL, e ElemType) int {
	locate := 0
	if L[SSLMAXSIZE-1].cur == 0 {
		fmt.Print("表空")
		return 0
	}
	k := L[SSLMAXSIZE-1].cur
	for k != 0 {
		if L[k].data == e {
			locate = k
			break
		}
		k = L[k].cur
	}
	return locate
}

// SSLEmpty 判断链表非空
func SSLEmpty(L *SSL) bool {
	k := L[SSLMAXSIZE-1].cur
	if k == 0 {
		return true
	}
	return false
}
