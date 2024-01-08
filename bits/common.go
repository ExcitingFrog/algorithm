package bits

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
