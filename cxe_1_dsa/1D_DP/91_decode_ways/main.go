package main

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	if len(s) < 2 {
		return 1
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}

func numDecodingsV2(s string) int {
	memo := make([]int, len(s))
	for i := range memo {
		memo[i] = -1
	}
	var dp func(i int) int
	dp = func(i int) int {
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		count := dp(i + 1)
		if i+1 < len(s) {
			if s[i:i+2] >= "10" && s[i:i+2] <= "26" {
				count += dp(i + 2)
			}
		}
		memo[i] = count
		return count
	}
	return dp(0)
}

func numDecodingsOptimal(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	if len(s) < 2 {
		return 1
	}
	prev := 1
	curr := 1
	for i := 2; i <= len(s); i++ {
		next := 0
		if s[i-1] != '0' {
			next += curr
		}
		if s[i-2:i] >= "10" && s[i-2:i] <= "26" {
			next += prev
		}
		prev, curr = curr, next
	}
	return curr
}
func main() {
	s := "226"
	println(numDecodings(s))
	println(numDecodingsV2(s))
	println(numDecodingsOptimal(s))
}