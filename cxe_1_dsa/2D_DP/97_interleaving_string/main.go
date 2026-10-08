package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)

	if m+n != len(s3) {
		return false
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			prev := dp[j]

			dp[j] = false

			if i > 0 && s1[i-1] == s3[i+j-1] {
				dp[j] = prev
			}

			if j > 0 && s2[j-1] == s3[i+j-1] {
				dp[j] = dp[j] || dp[j-1]
			}
		}
	}

	return dp[n]
}

func isInterleaveVer2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	memo := make([][]int, len(s1)+1)

	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i+j == len(s3) {
			memo[i][j] = 1
			return memo[i][j]
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if i == len(s1) {
			if s2[j] == s3[i+j] {
				memo[i][j] = dfs(i, j+1)
				return memo[i][j]
			}
			memo[i][j] = 0
			return memo[i][j]
		}

		if j == len(s2) {
			if s1[i] == s3[i+j] {
				memo[i][j] = dfs(i+1, j)
				return memo[i][j]
			}
			memo[i][j] = 0
			return memo[i][j]
		}
		result := 0
		if s1[i] == s3[i+j] {
			result = dfs(i+1, j)
		}
		if result == 0 && s2[j] == s3[i+j] {
			result = dfs(i, j+1)
		}
		memo[i][j] = result
		return result
	}

	return dfs(0, 0) == 1
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbcbcac"
	println(isInterleave(s1, s2, s3))
	println(isInterleaveVer2(s1, s2, s3))
}