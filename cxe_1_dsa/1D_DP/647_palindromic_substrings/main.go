package main

import (
	"fmt"
)

func countSubstrings(s string) int {
	n := len(s)
	dp := make([]bool, n)
	count := 0

	for i := n - 1; i >= 0; i-- {
		prev := false

		for j := i; j < n; j++ {
			temp := dp[j]

			if s[i] == s[j] {
				if j-i <= 2 || prev {
					dp[j] = true
					count++
				} else {
					dp[j] = false
				}
			} else {
				dp[j] = false
			}

			prev = temp
		}
	}

	return count
}

func countSubstringsV2(s string) int {
	count := 0
	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)
		expand(i, i+1)
	}

	return count
}

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
	fmt.Println(countSubstringsV2(s))
}