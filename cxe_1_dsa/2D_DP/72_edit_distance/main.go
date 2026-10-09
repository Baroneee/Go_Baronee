package main

import "fmt"
func minDistance(w1 string, w2 string) int {
    dp := make([]int, len(w1) + 1)
    for i := range dp {
        dp[i] = len(w1) - i
    }
    for j := len(w2) - 1; j >= 0; j-- {
        prev := dp[len(w1)]
        dp[len(w1)] += 1
        for i := len(w1) - 1; i >= 0; i-- {
            temp := dp[i]
            if w1[i] == w2[j] {
                dp[i] = prev
            } else {
                dp[i] = min(dp[i + 1] + 1, temp + 1, prev + 1)
            }
            prev = temp
        }
    }
    return dp[0]
}

func minDistanceVer2(w1 string, w2 string) int {
    memo := make([][]int, len(w1))
    for i := range memo {
        memo[i] = make([]int, len(w2))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i, j int) int 
    dfs = func(i, j int) int {
        if i == len(w1) && j == len(w2) {
            return 0
        }
        if i == len(w1) {
            return len(w2) - j
        }
        if j == len(w2) {
            return len(w1) - i
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }

        if w1[i] == w2[j] {
            memo[i][j] = dfs(i+1,j+1)
        } else {
            memo[i][j] = min(dfs(i+1,j)+1, dfs(i,j+1)+1, dfs(i+1,j+1)+1)
        }
        return memo[i][j]
    }
    return dfs(0,0)
}

func min(a, b, c int) int {
	if a < b {
		return a
	}

	if a < c {
		return a
	}

	if b < c {
		return b
	}
	return c
}

func main() {
	w1 := "horse"
	w2 := "ros"
	fmt.Println(minDistance(w1, w2))
}