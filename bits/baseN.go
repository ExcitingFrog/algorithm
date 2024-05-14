package bits

func baseN(n int, base int) string {
	res := ""
	if n == 0 {
		return "0"
	}

	for n != 0 {
		r := n % base
		if r < 0 {
			r -= base
			n -= r
		}
		if r >= 10 {
			res = string(r-10+'A') + res
		} else {
			res = string(r+'0') + res
		}
		n /= base
	}
	return res
}
