package main

import (
	"fmt"
	"sort"
)

func helper(candidates []int, target, cur int, list []int, ret *[][]int, deMap map[int]struct{}, start int) {
	if cur == target {
		r := make([]int, len(list))
		copy(r, list)
		*ret = append(*ret, r)
		return
	}

	for i := start; i < len(candidates) && cur+candidates[i] <= target; i++ {
		list = append(list, candidates[i])
		helper(candidates, target, cur+candidates[i], list, ret, deMap, i)
		list = list[:len(list)-1]
	}

	return
}

func combinationSum(candidates []int, target int) [][]int {
	var list []int
	var ret [][]int
	var deMap = make(map[int]struct{}, 0)
	sort.Ints(candidates)
	helper(candidates, target, 0, list, &ret, deMap, 0)
	return ret
}

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println("###", combinationSum(candidates, target))
}
