package main

import "fmt"

func longestPalindrome(s string) string {
	var start, end int = 0, 0

	if len(s) == 0 {
		return ""
	}

	n := len(s)
	ss := []rune(s)
	dp := make([][]bool, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]bool, n)
	}

	for i:=0; i<n-1; i++ {
		dp[i][i] = true
		if ss[i] == ss[i+1] {
			dp[i][i+1] = true
			start = i
			end = i+1
		} else {
			dp[i][i+1] = false
		}
	}

	for d := 2; d<n; d++ {
		for i:=0; i+d<n; i++ {
			if ss[i] == ss[i+d] && dp[i+1][i+d-1] {
				dp[i][i+d] = true
				if d+1 > end -start + 1 {
					start = i
					end = i + d
				}

			} else {
				dp[i][i+d] = false
			}
		}
	}


	return string(ss[start:end+1])
}

func main() {
	input := "ccc"
	fmt.Println(longestPalindrome(input))
}
