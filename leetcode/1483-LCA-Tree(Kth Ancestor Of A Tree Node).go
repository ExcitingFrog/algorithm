package leetcode

import "math/bits"

type TreeAncestor [][]int

func ConstructorTreeAncestor(n int, parent []int) TreeAncestor {
	tree := make([][]int, n)
	m := bits.Len(uint(n))
	for i, p := range parent {
		tree[i] = make([]int, m)
		tree[i][0] = p
	}
	for i := 0; i < m-1; i++ {
		for x := 0; x < n; x++ {
			if p := tree[x][i]; p != -1 {
				tree[x][i+1] = tree[p][i]
			} else {
				tree[x][i+1] = -1
			}
		}
	}
	return tree
}

func (t TreeAncestor) GetKthAncestor(node int, k int) int {
	n := bits.Len(uint(k))
	for i := 0; i < n; i++ {
		if k>>i&1 > 0 {
			node = t[node][i]
			if node < 0 {
				break
			}
		}
	}
	return node
}
