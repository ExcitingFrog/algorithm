package trees

import "fmt"

type SegmentTree struct {
	Nums  []int
	Trees []int
}

// 下标-二进制-存储区间和-前缀和计算
// 1 - 0001 - [1,1] - [1,1]
// 2 - 0010 - [1,2] - [1,2]
// 3 - 0011 - [3,3] - [1,2]+[3,3]
// 4 - 0100 - [1,4] - [1,4]
// 5 - 0101 - [5,5] - [1,4]+[5,5]
// 6 - 0110 - [5,6] - [1,4]+[5,6]
// 7 - 0111 - [7,7] - [1,4]+[5,6]+[7,7]
// 8 - 1000 - [1,8] - [1,8]
//   1	   2     3     4     5     6     7     8
// [1,1],[1,2],[3,3],[1,4],[5,5],[5,6],[7,7],[1,8]
// [ 1     3     3     10    5     11    7    36]

func BuildTree(nums []int) *SegmentTree {
	n := len(nums)
	trees := make([]int, n+1)
	for i, x := range nums {
		i++
		trees[i] += x
		nxt := i + (i & -i) //! core
		fmt.Println(i, nxt)
		// 1 2
		// 2 4
		// 3 4
		// 4 8
		// 5 6
		// 6 8
		// 7 8
		// 8 16
		if nxt <= n {
			trees[nxt] += trees[i]
		}
	}
	return &SegmentTree{nums, trees}
}

func (s *SegmentTree) Update(index, value int) {
	delta := value - s.Nums[index]
	s.Nums[index] = value
	for i := index + 1; i < len(s.Trees); i += i & -i { //! core
		s.Trees[i] += delta
	}
}

func (s *SegmentTree) PrefixSum(i int) (sum int) {
	for ; i > 0; i -= i & -i {
		sum += s.Trees[i]
	}
	return
}

func (s *SegmentTree) RangeSum(i, j int) int {
	fmt.Println(s.Trees)
	// [0 1 3 3 10 5 11 7 36]
	return s.PrefixSum(j+1) - s.PrefixSum(i)
}

func Test() {
	s := BuildTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(s.RangeSum(1, 3)) // 9
}
