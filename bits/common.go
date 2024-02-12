package bits

// refer: resources/html/分享｜从集合论到位运算，常见位运算技巧分类总结.html
func Common() {
	// cal musk
	nums := []int{1, 1, 0, 1}
	n := len(nums)
	var msk int
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 1 {
			msk += 1 << (n - i - 1)
			break
		}
	}

}

// foreach
// func Foreach(msk int) {
// 	for i := range 10 {
// 		if msk>>i&1 == 1 {
// 			// contain
// 		}
// 	}
// }
