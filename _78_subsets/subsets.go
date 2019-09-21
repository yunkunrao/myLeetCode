package main

import "fmt"

func subsets(nums []int) [][]int {

	var ret [][]int

	total := 1 << uint(len(nums))
	for i:=0; i < total; i++ {
		var tmp []int
		for j:=0; j < len(nums); j++ {
			if (i & (1 << uint(j))) > 0 {
				tmp = append(tmp, nums[j])
			}
		}
		ret = append(ret, tmp)
	}

	return ret
}

func main() {

	fmt.Println(subsets([]int{1,2,3}))
}