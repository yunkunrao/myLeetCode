package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var n = len(nums)
	var ret = make([]int, n)
	var tmp int = 1

	for i:=0; i<n; i++ {
		if i - 1 >= 0 {
			ret[i] = nums[i] * ret[i-1]
		} else {
			ret[i] = nums[i]
		}
	}

	for i:=n-1; i>=1; i-- {
		ret[i] = ret[i-1]
	}
	ret[0] = 1

	for i:=n-1; i>=0; i-- {
		ret[i] *= tmp
		tmp *= nums[i]
	}

	return ret
}

func main() {
	input := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(input))
}