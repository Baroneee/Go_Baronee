package main

import (
	"fmt"
	"strings"
)

// Version 1: Dynamic Programming (Optimized - Bottom-up)
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := range s {
		if !dp[i] {
			continue
		}
		for _, w := range wordDict {
			if strings.HasPrefix(s[i:], w) {
				dp[i+len(w)] = true
			}
		}
	}
	return dp[len(s)]
}

// Version 2: Brute Force (Decision Tree / DFS - Top-down)
func wordBreak_ver2(s string, wordDict []string) bool {

	var backtrack func(start int) bool
	backtrack = func(start int) bool {
		if start == len(s) {
			return true
		}

		for _, w := range wordDict {
			if strings.HasPrefix(s[start:], w) {
				if backtrack(start + len(w)) {
					return true 
				}
			}
		}

		return false
	}

	return backtrack(0)
}

func wordBreak_TopDown(s string, wordDict []string) bool {
    memo := make([]int, len(s)+1)
    var dp func(s string) bool
    dp = func(s string) bool {
        if len(s) == 0 {
            return true
        }

        n := len(s)
        if memo[n] != 0 {
            return memo[n] == 1
        }
        for _, w := range wordDict {
            if strings.HasPrefix(s,w) {
                if dp(s[len(w):]) {
                    memo[n] = 1
                    return true
                }
            }
        }
        memo[n] = -1
        return false
    }

    return dp(s)
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Println("Result (DP):", wordBreak(s, wordDict))
	fmt.Println("Result (DFS):", wordBreak_ver2(s, wordDict))
	fmt.Println("Result (Top-Down):", wordBreak_TopDown(s, wordDict))
}