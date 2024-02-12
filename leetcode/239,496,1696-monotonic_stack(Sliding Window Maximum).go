package leetcode

// 239
func MaxSlidingWindow(nums []int, k int) []int {
	sk := []int{}
	res := []int{}
	for i, x := range nums {
		// in
		for len(sk) > 0 && nums[sk[len(sk)-1]] <= x {
			sk = sk[:len(sk)-1]
		}
		sk = append(sk, i)
		// out
		if i-sk[0] >= k {
			sk = sk[1:]
		}
		// res
		if i >= k-1 {
			res = append(res, nums[sk[0]])
		}
	}
	return res
}

// 496
func NextGreaterElement(nums1 []int, nums2 []int) []int {
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

// 1696
func MaxResult(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	q := []int{0}
	for i := 1; i < n; i++ {
		if q[0] < i-k {
			q = q[1:]
		}
		f[i] = f[q[0]] + nums[i]
		for len(q) > 0 && f[q[len(q)-1]] <= f[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n-1]
}
