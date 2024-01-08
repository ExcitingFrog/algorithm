package leetcode

import "sort"

func LengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	res := 0
	for i, num := range nums {
		for j, x := range nums[:i] {
			if x < num {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
		res = max(res, f[i])
	}
	return res
}

func LengthOfLIS2(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	res := 0
	for i, num := range nums {
		for j, x := range nums[:i] {
			if x < num {
				f[i] = max(f[i], f[j])
			}
		}
		f[i]++
		res = max(res, f[i])
	}
	return res
}

// bilibili.com/video/BV1ub411Q7sB
func LengthOfLIS3(nums []int) int {
	g := nums[:0] // 原地修改
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}
