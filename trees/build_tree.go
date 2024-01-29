package trees

func BuildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	head := &TreeNode{}
	num := preorder[0]
	head.Val = num
	index := 0
	for i := range inorder {
		if inorder[i] == num {
			index = i
			break
		}
	}
	head.Left = BuildTreePre(preorder[1:index+1], inorder[:index])
	head.Right = BuildTreePre(preorder[index+1:], inorder[index+1:])
	return head
}

func BuildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	head := &TreeNode{}
	num := postorder[len(postorder)-1]
	head.Val = num
	index := 0
	for i := range inorder {
		if inorder[i] == num {
			index = i
			break
		}
	}
	head.Left = BuildTreePost(inorder[:index], postorder[:index])
	head.Right = BuildTreePost(inorder[index+1:], postorder[index:len(postorder)-1])
	return head
}
