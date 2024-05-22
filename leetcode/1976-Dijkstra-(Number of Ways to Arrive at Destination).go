package leetcode

import "math"

func CountPaths(n int, roads [][]int) int {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = math.MaxInt / 2
		}
	}
	for _, road := range roads {
		u, v, w := road[0], road[1], road[2]
		g[u][v] = w
		g[v][u] = w
	}
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt / 2
	}

	f := make([]int, n)
	f[0] = 1
	done := make([]bool, n)
	for {
		x := -1
		for i, ok := range done { //found not done and min dis
			if !ok && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x == n-1 { // return possible paths
			return f[n-1]
		}
		done[x] = true
		for y, d := range g[x] { // y , x->y:d
			newDis := dis[x] + d
			if newDis < dis[y] {
				dis[y] = newDis
				f[y] = f[x]
			} else if newDis == dis[y] {
				f[y] = (f[y] + f[x]) % (1e9 + 7)
			}
		}
	}
}
