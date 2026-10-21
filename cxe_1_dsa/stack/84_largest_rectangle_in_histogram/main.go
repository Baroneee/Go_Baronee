package main

func largestRectangleArea(heights []int) int {
	ans := 0
	s := []int{}
	for i := 0; i <= len(heights); i++ {
		h := 0
		if i < len(heights) {
			h = heights[i]
		}
		for len(s) > 0 && heights[s[len(s)-1]] >= h {
			height := heights[s[len(s)-1]]
			s = s[:len(s)-1]
			width := 0
			if len(s) == 0 {
				width = i
			} else {
				width = i - s[len(s)-1] - 1
			}
			ans = max(ans, width*height)
		}
		if i < len(heights) {
			s = append(s, i)
		}
	}
	return ans
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	println(largestRectangleArea(heights))
}
