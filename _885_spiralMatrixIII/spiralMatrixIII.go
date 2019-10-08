package main

import "fmt"

func getCircle(r0, c0, round int) [][]int {
	var ret [][]int
	if round < 0 {
		return nil
	}
	if round == 0 {
		ret = append(ret, []int{r0, c0})
		return ret
	}
	startI := r0 - (round - 1)
	startJ := c0 + round

	minI := r0 - round
	maxI := r0 + round
	minJ := c0 - round
	maxJ := c0 + round

	for i := startI; i <= maxI; i++ {
		ret = append(ret, []int{i, startJ})
	}

	for j := startJ - 1; j >= minJ+1; j-- {
		ret = append(ret, []int{maxI, j})
	}

	for i := maxI; i >= minI; i-- {
		ret = append(ret, []int{i, minJ})
	}

	for j := minJ + 1; j <= maxJ; j++ {
		ret = append(ret, []int{minI, j})
	}
	return ret
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	var ret [][]int
	var maxDistance int = -1

	left := c0
	right := C-1 - left
	top := r0
	bottom := R-1 - top

	if maxDistance < left {
		maxDistance = left
	}
	if maxDistance < right {
		maxDistance = right
	}
	if maxDistance < top {
		maxDistance = top
	}
	if maxDistance < bottom {
		maxDistance = bottom
	}

	for r := 0; r <= maxDistance; r++ {
		tmp := getCircle(r0, c0, r)
		for _, t := range tmp {
			if t[0] >= 0 && t[0] < R && t[1] >= 0 && t[1] < C {
				ret = append(ret, t)
			}
		}
	}

	return ret
}

func main() {

	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
	//fmt.Println(getCircle(0,0,2))
}
