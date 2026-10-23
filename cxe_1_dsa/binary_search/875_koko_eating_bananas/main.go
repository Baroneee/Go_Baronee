package main

func minEatingSpeed(piles []int, h int) int {
	end := 0
	for _, p := range piles {
		end = max(p, end)
	}
	l := 1
	r := end
	for l < r {
		k := l + (r-l)/2
		hours := 0
		for _, pile := range piles {
			hours += (pile + k - 1) / k
		}

		if hours <= h {
			r = k
		} else {
			l = k + 1
		}
	}
	return l
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	result := minEatingSpeed(piles, h)
	println("Minimum eating speed:", result)
}