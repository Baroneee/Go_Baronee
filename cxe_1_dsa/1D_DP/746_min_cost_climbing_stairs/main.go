package main

import "fmt"


func minCostClimbingStairsVer1(cost []int) int {
	n := len(cost)	
	if n < 2 {
		return 0
	}	
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func minCostClimbingStairsVer2(cost []int) int {
	n := len(cost)	
	if n < 2 {
		return 0
	}	
	prev, curr := 0, 0
	for i := 2; i <= n; i++ {
		temp := curr
		curr = min(curr+cost[i-1], prev+cost[i-2])
		prev = temp
	}
	return curr
}

func main() {
	cost := []int{10, 15, 20}
	
	fmt.Println("Result Ver 1:", minCostClimbingStairsVer1(cost))
	fmt.Println("Result Ver 2:", minCostClimbingStairsVer2(cost))
}
