package search

import "fmt"

// bilibili.com/video/BV1d54y1q7k7
func Search() {
	nums := []int{0, 1, 2, 5, 5, 5, 6, 7, 8}
	n := len(nums)
	left, right := -1, n
	for left+1 != right {
		mid := (left + right) >> 1
		if nums[mid] < 5 {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Println(left, right) // 2,3
	left, right = -1, n
	for left+1 != right {
		mid := (left + right) >> 1
		if nums[mid] <= 5 {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Println(left, right) // 5,6
}
