package main

import (
	"fmt"

	hr "github.com/Baroneee/Go_Baronee/cxe_1_dsa/1D_DP/house_robber"
)
	
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	var dp func(i int) int
	dp = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}
		if memo[i] != -1 {
			return memo[i]
		}
		memo[i] = max(dp(i-1), dp(i-2)+nums[i])
		return memo[i]
	}
	return dp(len(nums) - 1)
}

func robV2(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}


func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
	fmt.Println(robV2(nums))
	fmt.Println(hr.RobOptimal(nums))
}
