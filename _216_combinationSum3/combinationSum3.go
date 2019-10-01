package main

import "fmt"

func traceback(k, n, sum, begin int, ret *[][]int, list []int) {
	if k == 0 {
		if n == sum {
			tmp := make([]int, len(list))
			copy(tmp, list)
			*ret = append(*ret, tmp)
		}
		return
	}
	for i:=begin; i<10; i++ {
		list = append(list, i)
		traceback(k-1, n, sum+i, i+1, ret, list)
		list = list[:len(list)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	var ret = make([][]int,0)
	var list []int

	traceback(k, n, 0, 1, &ret, list)

	return ret
}

func main() {
	fmt.Println(combinationSum3(3, 7))
}
