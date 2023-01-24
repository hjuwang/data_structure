package avl

// 定义常量
const LH, EH, RH = 1, 0, -1

// RRotate 右旋操作
func RRotate(P *BiTree) {
	var L BiTree
	L = (*P).lchild
	//L的右子节点，挂在p的左子树
	(*P).lchild = L.rchild
	L.rchild = *P
	*P = L
}

// LRotate 左旋操作
func LRotate(P *BiTree) {
	var R BiTree
	//左旋挂在原来的右子树
	R = (*P).rchild
	(*P).rchild = R.lchild
	R.lchild = *P
	*P = R
}

// LeftBalance 左平衡旋转处理代码：（节点插在根节点的左孩子）
func LeftBalance(T *BiTree) {

	var L, Lr BiTree
	L = (*T).lchild
	switch L.bf {
	//新节点插在T的左孩子的左子树（直接右转根节点）LL(单右旋)
	case LH: //T左高，T.lchild 左高，单右旋T以后他们两个都平
		(*T).bf, L.bf = EH, EH
		RRotate(T)
		break
	//新节点插在T的左孩子的右子树	LR
	case RH: //T左高，T.lchiled右高，
		//提取左子树的右子树Lr
		Lr = L.rchild
		switch Lr.bf {
		//更新平衡因子
		case LH:
			(*T).bf = RH
			Lr.bf = EH
			break
		case EH:
			(*T).bf, L.bf = EH, EH
			break
		case RH:
			(*T).bf = EH
			L.bf = LH
			break
		}
		Lr.bf = EH
		//左旋T的左孩子
		LRotate((*BiTree)(&((*T).lchild)))
		//右旋根节点
		RRotate(T)
	}

}

// RightBalance 右平衡
func RightBalance(T *BiTree) {
	var R, Rl BiTree
	R = (*T).rchild
	switch R.bf {
	case RH:
		(*T).bf, R.bf = EH, EH
		LRotate(T)
		break
	case LH:
		Rl = R.lchild
		switch Rl.bf {
		case RH:
			(*T).bf = LH
			R.bf = EH
			break
		case EH:
			(*T).bf, R.bf = EH, EH
			break
		case LH:
			R.bf = RH
			break
		}
		Rl.bf = EH
		RRotate((*BiTree)(&(*T).rchild))
		LRotate(T)
	}

}
