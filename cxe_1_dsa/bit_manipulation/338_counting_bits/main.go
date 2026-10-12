package main

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ans[i] = ans[i/2] + i%2
	}
	return ans
}

func main() {
	n := 5 // Example input
	result := countBits(n)
	println(result) // Output: [0 1 1 2 1 2]
}
