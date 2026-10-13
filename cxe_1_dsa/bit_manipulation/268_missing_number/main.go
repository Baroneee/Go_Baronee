package main

func missingNumber(nums []int) int {
	ans := len(nums)
	for i := 0; i < len(nums); i++ {
		ans ^= i ^ nums[i]
	}
	return ans
}

func main() {
	nums := []int{3, 0, 1}
	result := missingNumber(nums)
	println(result) // Output: 2
}
