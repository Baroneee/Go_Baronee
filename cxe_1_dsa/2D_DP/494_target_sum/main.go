package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if abs(target) > sum {
		return 0
	}
	dp := make([]int, 2*sum+1)
	dp[sum] = 1
	for _, n := range nums {
		next := make([]int, 2*sum+1)
		for i := 0; i <= 2*sum; i++ {
			if dp[i] == 0 {
				continue
			}

			next[i-n] += dp[i]
			next[i+n] += dp[i]
		}

		dp = next
	}
	return dp[target+sum]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}