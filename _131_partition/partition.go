package main

import "fmt"

func isHuiWen(s []rune) bool {
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func helper(s []rune, ret *[][]string, cur int, curList *[]string, memo [][]int) {
	if len(s) == 0 {
		return
	}
	if cur == len(s) {
		dc := make([]string, len(*curList))
		copy(dc, *curList)
		*ret = append(*ret, dc)
		return
	}

	for i := cur; i < len(s); i++ {
		if memo[cur][i] == 1 {
			*curList = append(*curList, string(s[cur:i+1]))
			helper(s, ret, i+1, curList, memo)
			*curList = (*curList)[:len(*curList)-1]
		}
	}

}

func partition(s string) [][]string {
	var ret [][]string
	var curList []string
	ss := []rune(s)

	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for step := 1; step < n; step++ {
		for i := 0; i+step < n; i++ {
			if ss[i] == ss[i+step] && (step == 1 || dp[i+1][i+step-1] == 1) {
				dp[i][i+step] = 1
			}
		}
	}

	helper(ss, &ret, 0, &curList, dp)
	return ret
}

func main() {
	fmt.Println(partition("efe"))
}
