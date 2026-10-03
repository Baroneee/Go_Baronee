package main

import "fmt"
func canPartitionVer1(nums []int) bool {
    sum := 0
    for _, n := range nums {
        sum += n
    }

    if sum%2 != 0 {
        return false
    }

    target := sum / 2

    dp := make([]bool, target+1)
    dp[0] = true

    for _, num := range nums {
        for j := target; j >= num; j-- {
            if dp[j-num] {
                dp[j] = true
            }
        }
        
        if dp[target] {
            return true
        }     
    }

    return dp[target]
}

func canPartitionVer2(nums []int) bool {
    sum := 0
    for _, n := range nums {
        sum += n
    }

    if sum%2 != 0 {
        return false
    }

    target := sum / 2

    memo := make([][]bool, len(nums)+1)
    visited := make([][]bool, len(nums)+1)

    for i := range memo {
        memo[i] = make([]bool, target+1)
        visited[i] = make([]bool, target+1)
    }

    var dp func(i, target int) bool

    dp = func(i, target int) bool {
        if target == 0 {
            return true
        }

        if i == len(nums) {
            return false
        }

        if nums[i] > target {
            return dp(i+1, target)
        }

        if visited[i][target] {
            return memo[i][target]
        }

        visited[i][target] = true

        if dp(i+1, target-nums[i]) {
            memo[i][target] = true
            return true
        }

        if dp(i+1, target) {
            memo[i][target] = true
            return true
        }

        memo[i][target] = false
        return false
    }

    return dp(0, target)
}

func main() {
	fmt.Println(canPartitionVer1([]int{1, 5, 11, 5}))
	fmt.Println(canPartitionVer1([]int{1, 2, 3, 5}))
	fmt.Println(canPartitionVer2([]int{1, 5, 11, 5}))
	fmt.Println(canPartitionVer2([]int{1, 2, 3, 5}))
}
	