package balancebintree

// InsertAVL 递归生成平衡二叉树
func InsertAVL(T *BiTree, e int, taller *bool) bool {
	//插入新节点
	if (*T) == nil {
		(*T) = new(BiTNode)
		(*T).data = e
		(*T).lchild, (*T).rchild = nil, nil
		(*T).bf = EH
		*taller = true
	} else {
		//数据已经存在
		if e == (*T).data {
			*taller = false
			return false
		}
		if e < (*T).data {
			//在左子树中插入
			if !InsertAVL((*BiTree)(&(*T).lchild), e, taller) {
				//插入失败
				return false
			}
			if *taller {
				//左子树长高
				//检查T的平衡度
				switch (*T).bf {
				case LH: //原本左子树高，左子树又长高，左平衡调节,高度不变
					LeftBalance(T)
					*taller = false
					break
				case EH: //原本平衡，左子树长高，T左高右第低，高度增加
					(*T).bf = LH
					*taller = true
					break
				case RH: //原本右子树高，左子树长高，T左，右平,高度不变
					(*T).bf = EH
					*taller = false
					break
				}
			}

		} else {
			//在右子树中插入
			if !InsertAVL((*BiTree)(&(*T).rchild), e, taller) {
				//插入失败
				return false
			}
			if *taller {
				//右子树长高
				//检查T的平衡度
				switch (*T).bf {
				case LH: //原本左高，右子树增高，T平，高度不变
					(*T).bf = EH
					*taller = false
					break
				case EH: //原本平，右子树增高，T右高，高度增加
					(*T).bf = RH
					*taller = true
					break
				case RH: //原本右高，右边子树增高，T右平衡调节，高度不变
					RightBalance(T)
					*taller = false
					break
				}
			}

		}
	}
	return true
}
