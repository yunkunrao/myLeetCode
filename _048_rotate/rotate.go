package main

import "fmt"

func helper(matrix [][]int, l, r int) {
	var sur []int

	if l >= r {
		return
	}

	// |---| <---
	// |   |
	// |---|
	for j:=l; j <=r; j++ {
		sur = append(sur, matrix[l][j])
	}

	// |---|
	// |   | <--
	// |---|
	for j:=l+1; j<=r-1; j++ {
		sur = append(sur, matrix[j][r])
	}

	// |---|
	// |   |
	// |---| <---
	for j:=r; j>=l; j-- {
		sur = append(sur, matrix[r][j])
	}

	// |---|
	// | <-|---
	// |---|
	for j:=r-1; j>=l+1; j-- {
		sur = append(sur, matrix[j][l])
	}

	part1 := sur[:len(sur)-(r-l)]
	part2 := sur[len(sur)-(r-l):]
	part2 = append(part2, part1...)
	sur = part2

	// rotate
	var index int = 0
	for i:=l; i<=r; i++ {
		matrix[l][i] = sur[index]
		index++
	}

	for j:=l+1; j<=r-1; j++ {
		matrix[j][r] = sur[index]
		index++
	}

	for j:=r; j>=l; j-- {
		matrix[r][j] = sur[index]
		index++
	}

	for j:=r-1; j>=l+1; j-- {
		matrix[j][l] = sur[index]
		index++
	}
	helper(matrix, l+1, r-1)
}

func rotate(matrix [][]int)  {
	helper(matrix, 0, len(matrix)-1)
}

func main() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15,14,12,16},
	}

	rotate(matrix)

	fmt.Println(matrix)
}