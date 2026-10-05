package house_robber

func RobOptimal(nums []int) int {
	prev := 0
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		temp := curr
		curr = max(curr, prev+nums[i])
		prev = temp
	}
	return curr
}