package main

import "fmt"

func generateBound(nextNum, startX, startY int, matrix [][]int) int {

	n := len(matrix)

	// [startX, startY], [startX, (n-1)-startY]
	// [(n-1)-startX, startY],[(n-1)-startX, (n-1)-startY]

	// ----- <----this
	// |    |
	// -----
	for j:=startY; j <= (n-1)-startY && nextNum <= n*n; j++ {
		matrix[startX][j] = nextNum
		nextNum++
	}

	// -----
	// |    | <----this
	// -----
	for i:=startX+1; i <= (n-1)-startX-1 && nextNum <= n*n; i++ {
		matrix[i][(n-1)-startY] = nextNum
		nextNum++
	}

	// -----
	// |    |
	// ----- <----this
	for j:=(n-1)-startY; j >= startY && nextNum <= n*n; j-- {
		matrix[(n-1)-startX][j] = nextNum
		nextNum++
	}

	// -----
	// |<--|----this
	// -----
	for i:=(n-1)-startX-1; i >= startX+1 && nextNum <= n*n; i-- {
		matrix[i][startY] = nextNum
		nextNum++
	}

	return nextNum
}

func generateMatrix(n int) [][]int {
	var matrix = make([][]int, n)
	for i:=0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	var nextNum = 1
	for j:=0; j <= (n-1)/2; j++ {
		nextNum = generateBound(nextNum, j, j, matrix)
	}

	return matrix
}

func main() {

	fmt.Println(generateMatrix(5))
}