package trees

type AVLTreeNode struct {
	Height int32        //节点高
	Value  int32        //节点值
	Left   *AVLTreeNode //节点左儿子
	Right  *AVLTreeNode //节点右儿子
}

// GetHeight 返回节点的高
func GetHeight(a *AVLTreeNode) int32 {
	if a == nil {
		return 0
	}
	return a.Height
}

// Max 获取2数中的较大值
func Max(x, y int32) int32 {
	if x < y {
		return y
	}
	return x
}

// FindMax 查找最大节点
func FindMax(a *AVLTreeNode) *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.Right != nil {
		return FindMax(a.Right)
	}
	return a
}

// FindMin 查找最小值
func FindMin(a *AVLTreeNode) *AVLTreeNode {
	if a == nil {
		return nil
	}
	if a.Left != nil {
		return FindMin(a.Left)
	}
	return a
}

// GetBF 获取节点左右子树高度差绝对值
// 将二叉树上节点的左子树高度减去右子树高度取绝对值(Balance Factor)
func (a *AVLTreeNode) GetBF() int32 {
	var lh, rh int32
	if a.Left != nil {
		lh = GetHeight(a.Left)
	}
	if a.Right != nil {
		rh = GetHeight(a.Right)
	}
	bf := lh - rh
	if bf < 0 {
		bf = bf * -1
	}
	return bf
}

// LeftRotation 左旋
// a为最小失衡子数的根节点
// 问题：在右子树插入右孩子导致AVL失衡
// return 新的平衡树的根节点
func LeftRotation(a *AVLTreeNode) *AVLTreeNode {
	tmpNode := a.Right
	a.Right = tmpNode.Left
	tmpNode.Left = a
	a.Height = Max(GetHeight(a.Left), GetHeight(a.Right)) + 1
	tmpNode.Height = Max(GetHeight(tmpNode.Left), GetHeight(tmpNode.Right)) + 1
	return tmpNode
}

// RightRotation 右旋
// a为最小失衡树的根节点
// 问题：在左子树上插入左孩子导致AVL树失衡
// return 新的平衡树的根节点
func RightRotation(a *AVLTreeNode) *AVLTreeNode {
	tmpNode := a.Left
	a.Left = tmpNode.Right
	tmpNode.Right = a
	a.Height = Max(GetHeight(a.Left), GetHeight(a.Right)) + 1
	tmpNode.Height = Max(GetHeight(tmpNode.Left), GetHeight(tmpNode.Right)) + 1
	return tmpNode
}

// RightLeftRotation  右左双旋转
// 问题：通常因为在右子树上插入左孩子导致AVL失衡
// 解发：先右旋后左旋调整
// return 新的平衡树根节点
func RightLeftRotation(a *AVLTreeNode) *AVLTreeNode {
	a.Right = RightRotation(a.Right)
	return LeftRotation(a)
}

// LeftRightRotation  左右双选择
// 问题：通常因为在左子树上插入右孩子导致AVL失衡
// 解发：先左旋后右旋调整
// return 新的平衡树根节点
func LeftRightRotation(a *AVLTreeNode) *AVLTreeNode {
	a.Right = LeftRotation(a.Right)
	return RightRotation(a)
}

// InsertAVL avl树插入操作
// 递归
// 1个节点的高度初始值为1
// 返回插入后的根节点
func InsertAVL(tree *AVLTreeNode, v int32) *AVLTreeNode {
	if tree == nil { //空树，插入第一个节点
		tree := new(AVLTreeNode)
		tree.Value = v
	} else if v < tree.Value { //待插入的值比单前节点值小，往左子树插入
		tree.Left = InsertAVL(tree.Left, v)
		if GetHeight(tree.Left)-GetHeight(tree.Right) == 2 { //插入新节点后avl树失衡
			if v < tree.Left.Value { //情况1：v插入到左子树的左孩子节点，只需要进行一次右旋转
				tree = RightRotation(tree)
			} else if v > tree.Left.Value { //情况2：v插入到左子树的右孩子节点，需要先左旋转再右旋转
				tree = LeftRightRotation(tree)
			}
		}
	} else if v > tree.Value {
		tree.Right = InsertAVL(tree.Right, v)
		if GetHeight(tree.Right)-GetHeight(tree.Left) == 2 { //在右子树插入新节点后avl树失衡
			if v > tree.Right.Value { //情况4：v插入到右子树的右孩子节点，只需要进行一次左旋转
				tree = LeftRotation(tree)
			} else if v < tree.Right.Value { //情况3：v插入到右子树的左孩子节点，需要先右旋转再左旋转
				tree = RightLeftRotation(tree)
			}
		}
	}
	tree.Height = Max(GetHeight(tree.Left), GetHeight(tree.Right)) + 1
	return tree
}

// DeleteAVL 删除avl树中值为v的节点
// 维护二叉排序树规则
// 平衡avl二叉树
func DeleteAVL(tree *AVLTreeNode, v int32) *AVLTreeNode {
	if tree == nil {
		return nil
	}
	if tree.Value == v { //tree为待删除的节点
		//删除后维护成新的二叉排序树
		if tree.Left != nil && tree.Right != nil { //若待删除节点同时存在左右子树
			if GetHeight(tree.Left) > GetHeight(tree.Right) { //左子树高于右子树，则取左子树最大值取代被删除节点
				tmp := FindMax(tree.Left)
				tree.Value = tmp.Value
				tree.Left = DeleteAVL(tree.Left, tmp.Value)
			} else { //右子树高于左子树，则取右子树最小值取代被删除节点
				tmp := FindMin(tree.Right)
				tree.Value = tmp.Value
				tree.Right = DeleteAVL(tree.Right, tmp.Value)
			}
		} else {
			if tree.Left != nil {
				tree = tree.Left
			}
			if tree.Right != nil {
				tree = tree.Right
			}
		}
		return tree
	} else if tree.Value < v { //待删除节点比当前节点大
		tree.Right = DeleteAVL(tree.Right, v)
		if GetHeight(tree.Left)-GetHeight(tree.Right) == 2 { //删除节点后avl树失衡
			if v < tree.Left.Value { //情况1：左子树的左孩子节点过高，只需要进行一次右旋转
				tree = RightRotation(tree)
			} else if v > tree.Left.Value { //情况2：左子树的右孩子节点过高，需要先左旋转再右旋转
				tree = LeftRightRotation(tree)
			}
		}
	} else if tree.Value > v {
		tree.Left = DeleteAVL(tree.Left, v)
		if GetHeight(tree.Right)-GetHeight(tree.Left) == 2 { //在右子树插入新节点后avl树失衡
			if v > tree.Right.Value { //情况4：右子树的右孩子节点过高，只需要进行一次左旋转
				tree = LeftRotation(tree)
			} else if v < tree.Right.Value { //情况3：右子树的左孩子节点过高，需要先右旋转再左旋转
				tree = RightLeftRotation(tree)
			}
		}
	}
	return tree
}
