package main

import (
	"fmt"
)

func isConflict(rows []int, curColumn, curRow int) bool {
	if rows == nil {
		return false
	}

	for c, r := range rows {
		if r == curRow || c == curColumn {
			return true
		}
		if curColumn-c == curRow-r ||
			curColumn-c == r-curRow {
			return true
		}
	}
	return false
}

func displayArray(array [][]int) {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", array[i][j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func resetArray(array [][]int) {
	n := len(array)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			array[i][j] = 0
		}
	}
}

func totalNQueens(n int) int {
	var total int
	var rows []int
	var array = make([][]int, n)
	for i := 0; i < n; i++ {
		array[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		rows = append(rows, i)
		array[i][0] = 1

		for j := 1; j < n; j++ {
			for r := 0; r < n; r++ {
				//if i == 3 && j == 1 && r == 1 {
				if i == 3 && j == 1 {
					fmt.Println("hello", rows)
					//panic("##")
				}

				if isConflict(rows, j, r) {
					if r < (n - 1) {
						continue
					} else {
						r = rows[len(rows)-1]
						rows = rows[:len(rows)-1]
						j--
						array[r][j] = 0
						continue
					}
				} else {
					rows = append(rows, r)
					if len(rows) >= 2 &&
						rows[0] == 3 &&
						rows[1] == 1 {
						fmt.Println("@@@", rows)
						panic("")
					}

					array[r][j] = 1
					break
				}
			}
		}
		if len(rows) == n {
			total++
			displayArray(array)

			//for _, rr := range rows {
			//	fmt.Printf("## %d ", rr)
			//}
			//fmt.Println("")
		}

		resetArray(array)
		rows = nil
	}

	return total
}

func main() {
	fmt.Println(totalNQueens(8))
}
