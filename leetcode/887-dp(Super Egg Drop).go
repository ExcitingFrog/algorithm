package leetcode

func superEggDrop(k int, n int) int {
	// k eggs n height
	if n == 1 {
		return 1
	}
	f := make([][]int, n+1) // times, eggs
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for i := 1; i <= k; i++ {
		f[1][i] = 1
	}
	res := -1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			f[i][j] = 1 + f[i-1][j-1] + f[i-1][j]
		}
		if f[i][k] >= n {
			res = i
			break
		}
	}
	return res
}
