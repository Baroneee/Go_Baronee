package main

import (
	"fmt"
)
func maxProduct(nums []int) int {
    currMin := nums[0]
    currMax := nums[0]
    ans := nums[0]
    for i := 1; i < len(nums); i++{
        if nums[i] < 0 {
            currMin, currMax = currMax, currMin
        }
        
        currMin = min(nums[i], currMin*nums[i])
        currMax = max(nums[i], currMax*nums[i])
        ans = max(currMax, ans)
    }
    return ans
}   

func main() {	
	nums := []int{2, 3, -2, -4}
	fmt.Println(maxProduct(nums))
}

