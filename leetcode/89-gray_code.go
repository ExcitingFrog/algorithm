package leetcode

// i ⊕ i/2
func GrayCode(n int) []int {
	res := []int{}
	for i := 0; i < 1<<n; i++ {
		res = append(res, i^(i>>1))
	}
	return res
}
