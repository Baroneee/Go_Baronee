package main

import (
	"fmt"

	hr "github.com/Baroneee/Go_Baronee/cxe_1_dsa/1D_DP/house_robber"
) 

func rob2(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	return max(hr.RobOptimal(nums[1:]), hr.RobOptimal(nums[:len(nums)-1]))
}
func main() {
	nums := []int{2,3,2}
	fmt.Println(rob2(nums))	
}