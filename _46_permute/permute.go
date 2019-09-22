package main

import "fmt"

func insertAt(array []int, index, value int) []int {
	var ret []int
	if index == 0 {
		ret = append([]int{value}, array...)
	} else if index == len(array) {
		ret = append(array, value)
	} else {
		arrayCopy1 := append(make([]int, 0), array[0:index]...)
		arrayCopy2 := append(make([]int, 0), array[index:]...)
		ret = append(append(arrayCopy1, value), arrayCopy2...)
	}
	return ret
}

func permute(nums []int) [][]int {
	var ret [][]int

	if len(nums) == 1 {
		ret = append(ret, nums)
		return ret
	}

	tmp := permute(nums[0:len(nums)-1])
	for i :=0; i < len(nums); i++ {
		for j :=0; j < len(tmp); j++ {
			ret = append(ret, insertAt(tmp[j], i, nums[len(nums)-1]))
		}
	}
	return ret
}

func main() {

	fmt.Println(permute([]int{1, 2, 3}))
}