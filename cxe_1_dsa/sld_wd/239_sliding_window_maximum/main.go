package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		return []int{}
	}

	ans := make([]int, 0, len(nums)-k+1)
	dq := make([]int, 0, len(nums))
	head := 0

	for r := 0; r < len(nums); r++ {
		if head < len(dq) && dq[head] <= r-k {
			head++
		}

		for len(dq) > head && nums[dq[len(dq)-1]] <= nums[r] {
			dq = dq[:len(dq)-1]
		}

		dq = append(dq, r)

		if r >= k-1 {
			ans = append(ans, nums[dq[head]])
		}
	}

	return ans
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) // Output: [3 3 5 5 6 7]
}
