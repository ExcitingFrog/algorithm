package utils

func myPow(x float64, n int) float64 {
	dfs := func(x float64, n int) float64 {
		res := 1.0
		for n > 0 {
			if n%2 == 1 {
				res *= x
			}
			x *= x
			n /= 2
		}
		return res
	}
	if n < 0 {
		return 1 / dfs(x, -n)
	}
	return dfs(x, n)
}
