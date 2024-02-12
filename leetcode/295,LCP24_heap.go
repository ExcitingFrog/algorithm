package leetcode

import (
	"container/heap"
	"sort"
)

// LCP24
type BasePQ struct {
	sort.IntSlice
}

func (pq *BasePQ) Push(x any) {
	pq.IntSlice = append(pq.IntSlice, x.(int))
}

func (pq *BasePQ) Pop() any {
	n := len(pq.IntSlice)
	x := pq.IntSlice[n-1]
	pq.IntSlice = pq.IntSlice[:n-1]
	return x
}

func (pq *BasePQ) Top() int {
	return pq.IntSlice[0]
}

type MinPQ struct {
	*BasePQ
}

func (pq *MinPQ) Less(i, j int) bool {
	return pq.BasePQ.Less(i, j)
}

type MaxPQ struct {
	*BasePQ
}

func (pq *MaxPQ) Less(i, j int) bool {
	return pq.BasePQ.Less(j, i)
}

func numsGame(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	lower, upper := &MaxPQ{&BasePQ{}}, &MinPQ{&BasePQ{}}
	lowerSum, upperSum := int64(0), int64(0)
	mod := int64(1e9 + 7)
	for i := 0; i < n; i++ {
		x := nums[i] - i
		if lower.Len() == 0 || lower.Top() >= x {
			lowerSum += int64(x)
			heap.Push(lower, x)
			if lower.Len() > upper.Len()+1 {
				upperSum += int64(lower.Top())
				heap.Push(upper, lower.Top())
				lowerSum -= int64(heap.Pop(lower).(int))
			}
		} else {
			upperSum += int64(x)
			heap.Push(upper, x)
			if lower.Len() < upper.Len() {
				lowerSum += int64(upper.Top())
				heap.Push(lower, upper.Top())
				upperSum -= int64(heap.Pop(upper).(int))
			}
		}
		if (i+1)%2 == 0 {
			res[i] = int((upperSum - lowerSum) % mod)
		} else {
			res[i] = int((upperSum - lowerSum + int64(lower.Top())) % mod)
		}
	}
	return res
}

// 295
type MedianFinder struct {
	minQ, maxQ hp
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func Constructor295() MedianFinder {
	minq, maxq := hp{}, hp{}
	heap.Init(&minq)
	heap.Init(&maxq)
	hp := MedianFinder{}
	return hp
}

func (m *MedianFinder) AddNum(num int) {
	minq, maxq := &m.minQ, &m.maxQ
	if minq.Len() == 0 || num <= -minq.IntSlice[0] {
		heap.Push(minq, -num)
		if maxq.Len()+1 < minq.Len() {
			heap.Push(maxq, -heap.Pop(minq).(int))
		}
	} else {
		heap.Push(maxq, num)
		if maxq.Len() > minq.Len() {
			heap.Push(minq, -heap.Pop(maxq).(int))
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	minq, maxq := m.minQ, m.maxQ
	if minq.Len() > maxq.Len() {
		return float64(-minq.IntSlice[0])
	} else {
		return float64(maxq.IntSlice[0]-minq.IntSlice[0]) / 2
	}
}
