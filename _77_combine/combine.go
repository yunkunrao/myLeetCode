package main

import "fmt"

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

func combine(n int, k int) [][]int {
	var ret [][]int
	if k == 1 {
		for i := 1; i <= n; i++ {
			tmp := []int{i}
			ret = append(ret, tmp)
		}
		return ret
	}

	if n == k {
		var tmp []int
		for i := 1; i <= n; i++ {
			tmp = append(tmp, i)
		}
		ret = append(ret, tmp)
		return ret
	}

	for _, line := range combine(n-1, k-1) {
		line = append(line, n)
		ret = append(ret, line)
	}

	for _, line := range combine(n-1, k) {
		ret = append(ret, line)
	}

	return ret
}

func main() {
	ret := combine(4, 2)
	fmt.Println(ret)
}
