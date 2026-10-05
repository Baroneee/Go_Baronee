package main

func lengthOfLIS(nums []int) int {
	dp := []int{}
	dp = append(dp, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
		} else {
			l := 0
			r := len(dp) - 1
			for l < r {
				mid := l + (r-l)/2
				if dp[mid] >= nums[i] {
					r = mid
				} else {
					l = mid + 1
				}
			}
			dp[l] = nums[i]
		}
	}

	return len(dp)
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	println(lengthOfLIS(nums))
}
