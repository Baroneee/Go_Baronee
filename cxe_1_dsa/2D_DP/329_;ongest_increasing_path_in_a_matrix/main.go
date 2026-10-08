package main

func longestIncreasingPath(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])

	dp := make([][]int, rows)
	for r := range rows {
		dp[r] = make([]int, cols)
	}

	var dfs func(r, c int) int

	dfs = func(r, c int) int {
		if dp[r][c] != 0 {
			return dp[r][c]
		}

		dp[r][c] = 1

		if r > 0 && matrix[r-1][c] > matrix[r][c] {
			dp[r][c] = max(dp[r][c], 1+dfs(r-1, c))
		}

		if r+1 < rows && matrix[r+1][c] > matrix[r][c] {
			dp[r][c] = max(dp[r][c], 1+dfs(r+1, c))
		}

		if c > 0 && matrix[r][c-1] > matrix[r][c] {
			dp[r][c] = max(dp[r][c], 1+dfs(r, c-1))
		}

		if c+1 < cols && matrix[r][c+1] > matrix[r][c] {
			dp[r][c] = max(dp[r][c], 1+dfs(r, c+1))
		}

		return dp[r][c]
	}

	ans := 0

	for r := range rows {
		for c := range cols {
			ans = max(ans, dfs(r, c))
		}
	}

	return ans
}

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	println(longestIncreasingPath(matrix))
}