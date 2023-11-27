package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	nxt := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nxt[nums2[i]] = -1
		} else {
			nxt[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	res := []int{}
	for _, v := range nums1 {
		res = append(res, nxt[v])
	}
	return res
}
