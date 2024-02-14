package trees

func Traversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		// res = append(res, node.Val) //pre_order
		dfs(node.Left)
		// res = append(res, node.Val) //in_order
		dfs(node.Right)
		res = append(res, node.Val) //post_order
	}
	dfs(root)
	return res
}

// StackTraversal
func StackPreTraversal(root *TreeNode) []int {
	res := []int{}
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1].Right
		stk = stk[:len(stk)-1]
	}
	return res
}

func StackInTraversal(root *TreeNode) []int {
	res := []int{}
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		res = append(res, root.Val)
		stk = stk[:len(stk)-1]
		root = root.Right
	}
	return res
}

func StackPostTraversal(root *TreeNode) []int {
	res := []int{}
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return res
}
