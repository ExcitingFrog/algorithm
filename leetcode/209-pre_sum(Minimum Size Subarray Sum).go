package leetcode

import (
	"math"
	"sort"
)

func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := range nums {
		preSum[i+1] += preSum[i] + nums[i]
	}

	minLen := math.MaxInt32
	for i := 0; i < n; i++ {
		j := sort.Search(n-i+1, func(k int) bool {
			return preSum[i+k]-preSum[i] >= target
		})
		if j != n-i+1 {
			minLen = min(minLen, j)
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
