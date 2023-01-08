package main

import "fmt"

func main() {

	//c := &LNode{13, nil}
	//b := &LNode{12, c}
	//a := &LNode{11, b}
	//head := &LNode{3, a}
	//var num ElemType
	//LinkListDelete(head, 3, &num)
	//fmt.Print(num)

	//var a [3]int = [3]int{1, 2, 3}
	//j := a
	//j[0] = 5
	//fmt.Print(a[0])

	var ssl SSL
	InitSSL(&ssl)
	//fmt.Print(ssl[SSLMAXSIZE-1].cur)
	SSLInsert(&ssl, 1, 11)
	SSLInsert(&ssl, 2, 12)

	//SSLInsert(&ssl, 2, 12)

	//插入两个会有问题

	var e ElemType
	SSLDelete(&ssl, 1, &e)
	fmt.Println(SSLLength(&ssl))
	fmt.Println(e)

}
