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

// Version 2: Brute Force (Decision Tree / BFS - Top-down)
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

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Println("Result (DP):", wordBreak(s, wordDict))
	fmt.Println("Result (DFS):", wordBreak_ver2(s, wordDict))
}