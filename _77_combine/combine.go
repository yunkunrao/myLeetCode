package main

import (
	"fmt"
	"sort"
)

func compareSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, ss := range s1 {
		if s2[i] != ss {
			return false
		}
	}
	return true
}

func findKSumPairs(array []int, k int) [][]int {
	var ret [][]int
	start := 0
	end := len(array) - 1

	sort.Ints(array)
	for start < end {
		if array[end] - array[start]> k {
			end--
		} else if array[end] - array[start] < k {
			start++
		} else {
			tmp := []int{array[start], array[end]}
			ret = append(ret, tmp)
			start++
			end--
		}
	}
	return ret
}

func main() {
	input := []int{1, 2, 3, 5, 7}
	// (1, 3) (3, 5) (5, 7)
	fmt.Println(findKSumPairs(input, 2))
}
