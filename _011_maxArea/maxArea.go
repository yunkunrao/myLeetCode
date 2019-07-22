package main

import "fmt"

func maxArea(height []int) int {

	var max = 0

	begin := 0
	end := len(height) - 1

	for begin < end {

		if height[begin] < height[end] {
			if max < height[begin]*(end-begin) {
				max = height[begin] * (end - begin)
			}

			begin++
		} else {
			if max < height[end]*(end-begin) {
				max = height[end] * (end - begin)
			}

			end--
		}

	}

	return max
}

func main() {
	var succeed = true

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	if 49 != maxArea(input) {
		succeed = false
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!", maxArea(input))
	}
}
