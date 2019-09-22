package main

import "fmt"

func generateAll (cur string, num, open, close int, ret *[]string) {
	if len(cur) == 2*num {
		*ret = append(*ret, cur)
		return
	}
	if open < num {
		generateAll(cur + "(", num, open + 1, close, ret)
	}
	if close < open {
		generateAll(cur + ")", num, open, close + 1,ret)
	}
}

func generateParenthesis(n int) []string {
	var ret []string

	generateAll("", n,0,0, &ret)

	return ret
}

func main() {

	fmt.Println(generateParenthesis(1))
}

