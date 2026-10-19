package main

import "fmt"

func minWindow(s string, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}

	freq := make([]int, 256)
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	need := len(t)
	l := 0
	minL := len(s) + 1
	ans := ""

	for r := 0; r < len(s); r++ {
		c := s[r]

		if freq[c] > 0 {
			need--
		}
		freq[c]--

		for need == 0 {
			if r-l+1 < minL {
				minL = r - l + 1
				ans = s[l : r+1]
			}

			c := s[l]
			freq[c]++
			if freq[c] > 0 {
				need++
			}
			l++
		}
	}

	return ans
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t)) // Output: "BANC"
}