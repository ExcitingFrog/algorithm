package search

// zhihu.com/question/21923021/answer/1032665486
// resources/html/如何更好地理解和掌握 KMP 算法
func Kmp(text, pattern string) (pos []int) {
	nxt := buildNx(pattern)
	tIndex, pIndex := 0, 0
	for tIndex < len(text) {
		if text[tIndex] == pattern[pIndex] {
			tIndex++
			pIndex++
		} else if pIndex != 0 {
			pIndex = nxt[pIndex-1]
		} else {
			tIndex++
		}

		if pIndex == len(pattern) {
			pos = append(pos, tIndex-pIndex)
			pIndex = nxt[pIndex-1]
		}
	}
	return
}

func buildNx(p string) []int {
	m := len(p)
	nxt := make([]int, m)
	i, now := 1, 0
	for i < m {
		if p[i] == p[now] {
			now++
			nxt[i] = now
			i++
		} else if now != 0 {
			now = nxt[now-1]
		} else {
			nxt[i] = now
			i++
		}
	}
	return nxt
}
