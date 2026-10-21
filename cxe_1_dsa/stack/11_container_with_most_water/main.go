package main

func maxArea(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		square := (r - l) * min(height[r], height[l])
		ans = max(square, ans)
		if height[r] >= height[l] {
			l++
			continue
		} else {
			r--
			continue
		}
	}
	return ans
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	println(maxArea(height))
}