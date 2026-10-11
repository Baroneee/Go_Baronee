package main

import "fmt"

func isMatch(s string, p string) bool {
    m, n := len(s), len(p)

    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }

    dp[m][n] = true

    for j := n - 2; j >= 0; j-- {
        if p[j+1] == '*' {
            dp[m][j] = dp[m][j+2]
        }
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            firstMatch := p[j] == '.' || s[i] == p[j]

            if j+1 < n && p[j+1] == '*' {
                dp[i][j] = dp[i][j+2] ||
                    (firstMatch && dp[i+1][j])
            } else {
                dp[i][j] = firstMatch && dp[i+1][j+1]
            }
        }
    }

    return dp[0][0]
}


func isMatchVer2(s string, p string) bool {
    m, n := len(s), len(p)

    memo := make([][]int, m+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if memo[i][j] != -1 {
            return memo[i][j] == 1
        }

        result := false

        if i == 0 && j == 0 {
            result = true
        } else if j == 0 {
            result = false
        } else if p[j-1] == '*' {
            result = dfs(i, j-2)

            if i > 0 && (p[j-2] == '.' || s[i-1] == p[j-2]) {
                result = result || dfs(i-1, j)
            }
        } else {
            if i > 0 && (p[j-1] == '.' || s[i-1] == p[j-1]) {
                result = dfs(i-1, j-1)
            }
        }

        if result {
            memo[i][j] = 1
        } else {
            memo[i][j] = 0
        }

        return result
    }

    return dfs(m, n)
}

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p)) // Output: true
	fmt.Println(isMatchVer2(s, p)) // Output: true
}