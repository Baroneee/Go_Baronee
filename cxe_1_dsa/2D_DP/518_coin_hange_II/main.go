package main

import "fmt"
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		c := coins[i-1]

		for a := 1; a <= amount; a++ {
			dp[i][a] = dp[i-1][a]

			if a >= c {
				dp[i][a] += dp[i][a-c]
			}
		}
	}

	return dp[n][amount]
}

func main() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}