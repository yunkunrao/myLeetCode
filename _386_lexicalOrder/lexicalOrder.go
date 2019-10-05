package main

import "fmt"

func helper(ret *[]int, cur, n int) {
	if cur > n {
		return
	}
	if cur != 0 {
		*ret = append(*ret, cur)
	}

	for i := 0; i <= 9; i++ {
		if cur == 0 {
			if i == 0 {
				continue
			}
		}
		helper(ret, 10*cur+i, n)
	}
}

func lexicalOrder(n int) []int {
	var ret []int
	helper(&ret, 0, n)
	return ret
}

func main() {
	fmt.Println(lexicalOrder(13))
}
