package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	var start = 0
	var end = len(nums)-1

	if len(nums) == 1 {
		return nums[0]
	}

	// 0 1 2 3 4 5 6 7 8
	// 1 1 3 3 2 4 4 8 8
	for start <= end {
		mid := (start+end)/2
		if mid %2 == 1 {
			if mid -1 >= 0 && nums[mid-1] == nums[mid] {
				start = mid + 1
			} else if mid + 1 <= end && nums[mid] == nums[mid+1] {
				end = mid - 1
			} else {
				return nums[mid]
			}
		} else {
			if mid+1 <= end && nums[mid] == nums[mid+1] {
				start = mid + 2
			} else if mid -1 >= 0 && nums[mid-1] == nums[mid] {
				end = mid - 2
			} else {
				return nums[mid]
			}
		}
	}
	return -1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(singleNonDuplicate(nums))
}