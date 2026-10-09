package main

import "fmt"
func numDistinct(s string, t string) int {
    if len(s) < len(t) {
        return 0
    }
    dp := make([]int, len(t)+1)
    for i := range dp {
        dp[i] = 0
        if i == len(t) {
            dp[i] = 1
        }
    }
    for i := len(s) - 1; i >= 0; i-- {
        for j := len(t) - 1; j >= 0; j-- {
            if s[i] == t[j] {
                dp[j] += dp[j+1]
            } 
        }
		fmt.Println(i, string(s[i]), dp)
    }
    return dp[0]
}

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}
