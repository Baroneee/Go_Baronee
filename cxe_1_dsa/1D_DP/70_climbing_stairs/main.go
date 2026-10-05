package main

import "fmt"

func climbStairs(n int) int {
	check := make([]int, n+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n == 1 {
			return 1
		}
		if n == 2 {
			return 2
		}
		if check[n] != 0 {
			return check[n]
		}
		check[n] = dp(n-1) + dp(n-2)
		return check[n]
	}
	return dp(n)
}

func climbStairsOptimal(n int) int {
	prev := 0
	curr := 1
	for i := 1; i <= n; i++ {
		temp := curr
		curr = prev + curr
		prev = temp
	}
	return curr
}

func main() {
	n := 5
	result := climbStairs(n)
	fmt.Println(result)
	fmt.Println(climbStairsOptimal(n))
}