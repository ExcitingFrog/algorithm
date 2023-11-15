package leetcode

func ReverseLink(root *LinkNode) *LinkNode {
	if root == nil || root.Next == nil {
		return root
	}
	last := ReverseLink(root.Next)
	root.Next.Next = root
	root.Next = nil
	return last
}
