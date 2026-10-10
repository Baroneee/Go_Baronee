package main

import "fmt"

func maxCoins(nums []int) int {
    n := len(nums)

    values := make([]int, n+2)
    values[0] = 1
    values[n+1] = 1
    copy(values[1:n+1], nums)

    memo := make([][]int, n+2)
    for i := range memo {
        memo[i] = make([]int, n+2)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(l, r int) int

    dfs = func(l, r int) int {
        if l > r {
            return 0
        }

        if memo[l][r] != -1 {
            return memo[l][r]
        }

        best := 0

        for k := l; k <= r; k++ {
            coins := values[l-1] * values[k] * values[r+1]

            left := dfs(l, k-1)
            right := dfs(k+1, r)

            total := left + right + coins

            if total > best {
                best = total
            }
        }

        memo[l][r] = best
        return best
    }

    return dfs(1, n)
}

func maxCoinsVer2(nums []int) int {
    n := len(nums)

    values := make([]int, n+2)
    values[0] = 1
    values[n+1] = 1

    copy(values[1:n+1], nums)

    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    for length := 1; length <= n; length++ {
        for l := 1; l <= n-length+1; l++ {
            r := l + length - 1
            best := 0
            for k := l; k <= r; k++ {
                coins := values[l-1] * values[k] * values[r+1]

                left := dp[l][k-1]
                right := dp[k+1][r]

                total := left + right + coins

                if total > best {
                    best = total
                }
            }
            dp[l][r] = best
        }
    }

    return dp[1][n]
}

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
	fmt.Println(maxCoinsVer2(nums))
}
