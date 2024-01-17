package leetcode

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

type CountIntervals struct {
	*redblacktree.Tree
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{
		redblacktree.NewWithIntComparator(),
		0,
	}
}

func (c *CountIntervals) Add(left int, right int) {
	for node, _ := c.Ceiling(left); node != nil && node.Value.(int) <= right; node, _ = c.Ceiling(left) {
		r, l := node.Key.(int), node.Value.(int)
		c.cnt -= (r - l + 1)
		if l <= left {
			left = l
		}
		if r >= right {
			right = r
		}
		c.Remove(node.Key)
	}
	c.Put(right, left)
	c.cnt += (right - left + 1)
}

func (c *CountIntervals) Count() int {
	return c.cnt
}
