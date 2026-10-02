package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = min(dp[i-c]+1, dp[i])
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func coinChangeRecursive(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	memo := make([]int, amount+1)
	for i := range memo {
		memo[i] = -2
	}
	var dp func(target int) int

	dp = func(target int) int {
		if target == 0 {
			return 0
		}

		if target < 0 {
			return -1
		}

		if memo[target] != -2 {
			return memo[target]
		}
		ans := -1
		for _, c := range coins {
			sub := dp(target - c)
			if sub != -1 {
				candidate := sub + 1
				if ans == -1 || candidate < ans {
					ans = candidate
				}
			}
		}
		memo[target] = ans
		return ans
	}
	return dp(amount)
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	result := coinChange(coins, amount)
	resultRecursive := coinChangeRecursive(coins, amount)
	if result == -1 {
		println("No solution")
	} else {
		println("Minimum coins needed:", result)
	}
	if resultRecursive == -1 {
		println("No solution (recursive)")
	} else {
		println("Minimum coins needed (recursive):", resultRecursive)
	}
}