package main

import (
	"fmt"
)
func longestCommonSubsequenceVer2(text1 string, text2 string) int {
    dp := make([]int, len(text1)+1)

    for i := 1; i <= len(text2); i++ {
        prev := 0 

        for j := 1; j <= len(text1); j++ {
            temp := dp[j] 

            if text1[j-1] == text2[i-1] {
                dp[j] = prev + 1
            } else {
                dp[j] = max(dp[j], dp[j-1])
            }

            prev = temp
        }
    }

    return dp[len(text1)]
}

func longestCommonSubsequence(text1 string, text2 string) int {
    memo := make([][]int, len(text1))
    for i := range memo {
        memo[i] = make([]int, len(text2))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dp func(i, j int) int
    dp = func(i, j int) int {
        if i == len(text1) || j == len(text2) {
            return 0 
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }

        if text1[i] == text2[j] {
            memo[i][j] = 1 + dp(i+1, j+1)
        } else {
            memo[i][j] = max(
                dp(i+1, j),
                dp(i, j+1),
            )
        }

        return memo[i][j]
    }

    return dp(0,0)
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2)) 
	fmt.Println(longestCommonSubsequenceVer2(text1, text2))
}