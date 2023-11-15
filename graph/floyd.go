package graph

// n cities, m edges
// edge = [u, v, w]  u, v: city, w: weight
func floyd(n int, edges [][]int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = 1e9
		}
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		g[u][v], g[v][u] = w, w
	}

	memo := make([][][]int, n) //i,j,k weight
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, n)
		}
	}

	var dfs func(k int, i int, j int) int
	dfs = func(k, i, j int) int {
		if k < 0 {
			return g[i][j]
		}
		if memo[i][j][k] != 0 {
			return memo[i][j][k]
		}

		res := min(dfs(k-1, i, j), dfs(k-1, i, k)+dfs(k-1, k, j))
		memo[i][j][k] = res
		return res
	}

	return
}
