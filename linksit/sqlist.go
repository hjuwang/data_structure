package main

// 定义顺序结构类型
type SQList struct {
	data   [MAXSIZE]ElemType
	length int //当前长度

}

// InitList 初始化线性表
func InitList(l *SQList) {
	//分配数组
	l.data = [MAXSIZE]ElemType{}
	l.length = 0

}

// ListEmpty 判断线性表是否为空
func ListEmpty(l *SQList) bool {
	if l.length == 0 {
		return true
	}
	return false
}

// ListLength 返回线性表元素的个数
func ListLength(l *SQList) int {
	return l.length
}

// ListInset 插入元素,位置i
func ListInset(l *SQList, i int, e ElemType) Status {
	if l.length == MAXSIZE {
		return Error
	}
	//超出插入范围
	if i < 1 || i > l.length+1 {
		return Error
	}
	//在下标i-1处插入数据
	if i < l.length {
		//
		for k := l.length - 1; k > i-1; k-- {
			//向后移动
			l.data[k+1] = l.data[k]
		}
	}
	l.data[i-1] = e
	l.length++
	return OK

}

// ListDelete 删除元素，位置i,并用e返回其值
func ListDelete(l *SQList, i int, e *ElemType) Status {
	//表空无法删除
	if l.length == 0 {
		return Error
	}
	//非法
	if i < 1 || i > l.length {
		return Error
	}
	*e = l.data[i-1]
	//i 属于[1,l.length-1]
	if i < l.length {
		for k := i; k < l.length; k++ {
			//注意第一次循环和最后一次循环即可
			l.data[k-1] = l.data[k]
		}

	}
	l.length--
	return OK
}

// ClearList 清空一个线性表
func ClearList(l *SQList) Status {
	end := l.length
	for i := 0; i < end; i++ {
		l.data[i] = 0
		l.length--
	}
	if l.length == 0 {
		return OK
	} else {
		return false
	}

}

// GetElem 返回线性表中的地i个元素
func GetElem(l *SQList, i int, e *ElemType) Status {
	if i < 1 || i > l.length {
		return Error
	}
	*e = l.data[i-1]
	return OK
}

// LocalElem 在线性表中查找与给定值，成功返回位置，不成功返回0
func LocalElem(l *SQList, e ElemType) int {
	index := 0
	for i := 0; i < l.length; i++ {
		if l.data[i] == e {
			index = i + 1
			break
		}
	}
	return index
}

//func main() {
//	//初始化
//	var list SQList
//	InitList(&list)
//	//插入元素
//	ListInset(&list, 1, 1)
//	ListInset(&list, 2, 2)
//	ListInset(&list, 3, 3)
//	ListInset(&list, 4, 4)
//	ListInset(&list, 5, 5)
//
//	//删除元素
//	var e ElemType
//	ListDelete(&list, 3, &e)
//	fmt.Print(e)
//	fmt.Print(ListEmpty(&list))
//	ClearList(&list)
//	fmt.Print(ListEmpty(&list))
//
//}

//
