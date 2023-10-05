package union

type UnionFind struct {
	count  int
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	un := &UnionFind{
		count: n,
	}
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	un.parent = parent
	un.size = size

	return un
}

func (u *UnionFind) Union(a, b int) {
	rootA := u.Find(a)
	rootB := u.Find(b)
	if rootA == rootB {
		return
	}
	if u.size[a] > u.size[b] {
		u.parent[b] = rootA
		u.size[a] += u.size[b]
	} else {
		u.parent[a] = rootB
		u.size[b] += u.size[a]
	}
	u.count--
}

func (u *UnionFind) Find(x int) int {
	if x != u.parent[x] {
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}
