package leetcode

func maxProduct(words []string) int {
	mask := make([]int, len(words))
	for i := range mask {
		for _, b := range words[i] {
			mask[i] |= 1 << (b - 'a')
		}
	}
	res := 0
	for i := range mask {
		for j := i + 1; j < len(mask); j++ {
			if mask[i]&mask[j] == 0 {
				if len(words[i])*len(words[j]) > res {
					res = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return res
}
