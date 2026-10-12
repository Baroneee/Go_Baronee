package main

func singleNumber(nums []int) int {
	ans := 0
	for _, n := range nums {
		ans ^= n
	}
	return ans
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	println(result) // Output: 4
}