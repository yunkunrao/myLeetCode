package main

import "fmt"

func columnReverse(A [][]int) {
	for j := 0; j < len(A[0]);j++ {
		ones := 0
		zeros := 0
		for i:=0; i < len(A); i++ {
			if A[i][j] == 1 {
				ones++
			} else {
				zeros++
			}
		}
		if ones < zeros {
			for k := 0; k < len(A); k++ {
				if A[k][j] == 1 {
					A[k][j] = 0
				} else {
					A[k][j] = 1
				}
			}
		}
	}
}

func matrixScore(A [][]int) int {
	for _, line := range A {
		if line[0] == 0 {
			for i, num := range line {
				if num == 0 {
					line[i] = 1
				} else {
					line[i] = 0
				}
			}
		}
	}

	columnReverse(A)

	var sum int = 0

	for _, line := range A {
		for i, num := range line {
			sum += num * (1 << uint(len(line) - i - 1))
		}
	}
	return sum
}

func main() {
	testData := [][]int{{1,1},{1,1},{0,1}}
	fmt.Println(matrixScore(testData))
}
