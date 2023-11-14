package utils

import (
	"fmt"
	"sort"
)

func Binary() {
	nums := []int{0, 1, 2, 5, 5, 5, 6, 7, 8}
	n := len(nums)
	// return the index of the first element that >= 5
	index1 := sort.Search(n, func(i int) bool {
		return nums[i] >= 5
	})
	// return the index of the first element that > 5
	index2 := sort.Search(n, func(i int) bool {
		return nums[i] > 5
	})
	// not found return len(nums)
	index3 := sort.Search(n, func(i int) bool {
		return nums[i] > 100
	})
	// return 0
	index4 := sort.Search(n, func(i int) bool {
		return nums[i] < 100
	})
	fmt.Println(index1, index2, index3, index4) // 3 6 9 0
}
