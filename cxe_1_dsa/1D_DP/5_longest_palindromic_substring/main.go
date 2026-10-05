package main

import (
	"fmt"
)
func longestPalindrome(s string) string {
    ans := ""
    for i := range s {
        even := extend(s, i, i+1)
        if len(even) > len(ans) {
            ans = even
        }

        odd := extend(s, i, i)
        if len(odd) > len(ans) {
            ans = odd
        }
    }
    return ans
}
func extend(s string, i, j int) string {
    for i >= 0 && j < len(s) && s[i] == s[j]  {
        i--
        j++
    }
    return s[i+1 : j]
}

func main() {
    s := "babad"
    fmt.Println(longestPalindrome(s))
}