package leetcode

func CountSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}
	for i := range matrix {
		for j := range matrix[i] {
			d[i+1][j+1] = d[i][j+1] + d[i+1][j] - d[i][j] + matrix[i][j]
		}
	}
	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for i1 := 0; i1 <= m-i; i1++ {
				for j1 := 0; j1 <= n-j; j1++ {
					if d[i+i1][j+j1]-d[i-1][j+j1]-d[i+i1][j-1]+d[i-1][j-1] == (i1+1)*(j1+1) {
						res++
					}
				}
			}
		}
	}

	return res
}

func NumSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
	}
	for i := range mat {
		for j := range mat[i] {
			d[i+1][j+1] = d[i][j+1] + d[i+1][j] - d[i][j] + mat[i][j]
		}
	}
	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			for i1 := 0; i1 <= m-i; i1++ {
				for j1 := 0; j1 <= n-j; j1++ {
					if d[i+i1][j+j1]-d[i-1][j+j1]-d[i+i1][j-1]+d[i-1][j-1] == (i1+1)*(j1+1) {
						res++
					}
				}
			}
		}
	}

	return res
}
