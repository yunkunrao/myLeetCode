package main

import "fmt"

func findDuplicates(nums []int) []int {
	var ret []int

	for _,num := range nums {
		if num < 0 {
			num *= -1
		}
		if nums[num-1] > 0 {
			nums[num-1] *= -1
		} else {
			ret = append(ret, num)
		}
	}

	return ret
}

func main() {
	nums := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(nums))
}