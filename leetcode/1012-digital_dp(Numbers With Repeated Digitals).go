package leetcode

import "strconv"

func NumDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	length := len(s)
	memo := make([][1 << 10]int, length)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var f func(i, msk int, isLimit, isNum bool) int
	f = func(i, msk int, isLimit, isNum bool) int {
		if i == length {
			if isNum {
				return 1
			}
			return 0
		}
		if !isLimit && isNum {
			if memo[i][msk] > 0 {
				return memo[i][msk]
			}
		}
		res := 0
		if !isNum {
			res += f(i+1, msk, false, false)
		}

		d := 0
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if msk>>d&1 == 0 {
				res += f(i+1, msk|1<<d, isLimit && d == up, true)
			}
		}
		return res
	}
	return n - f(0, 0, true, false)
}
