package main

import (
	"fmt"
	"sort"
)

func helper(people [][]int) [][]int {
	var minHArray = make([][]int, 0)
	var minH int = int(^uint(0) >> 1)

	if len(people) == 0 {
		return nil
	}

	for _, line := range people {
		if line[0] < minH {
			minH = line[0]
		}
	}

	peo := make([][]int, 0)
	for _, line := range people {
		if line[0] != minH {
			peo = append(peo, line)
		} else {
			minHArray = append(minHArray, line)
		}
	}

	sort.SliceStable(minHArray, func(i, j int) bool {
		return minHArray[i][1] < minHArray[j][1]
	})

	repeo := helper(peo)

	for _, line := range minHArray {
		repeo = append(repeo[:line[1]], append([][]int{line}, repeo[line[1]:]...)...)
	}

	return repeo
}

func reconstructQueue(people [][]int) [][]int {
	return helper(people)
}

func main() {
	people := [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}

	fmt.Println(reconstructQueue(people))
}
