package main

import "fmt"

func deepCopy(board [][]int) [][]int {
	var ret [][]int

	for _, b := range board {
		tmp := make([]int, len(b))
		copy(tmp, b)
		ret = append(ret, tmp)
	}
	return ret
}

func countSuroundOnes(board [][]int, i, j int) int {
	// (i-1, j-1) (i-1, j) (i-1, j+1)
	// (i, j-1)   (i, j)   (i, j+1)
	// (i+1, j-1) (i+1, j) (i+1, j+1)
	var total int = 0
	M := len(board)
	N := len(board[0])

	for m:=i-1; m<=i+1; m++ {
		if m < 0 {
			continue
		}
		if m > M-1 {
			break
		}
		for n:=j-1; n<=j+1; n++ {
			if n < 0 {
				continue
			}
			if n > N-1 {
				break
			}
			if m == i && n == j {
				continue
			}
			total += board[m][n]
		}
	}
	return total
}

func gameOfLife(board [][]int)  {
	// for live
	// [0, 2) dead
	// [2, 3] live
	// (3, infi) dead

	// for dead
	// [3] relive

	dc := deepCopy(board)

	for i, b := range board {
		for j, _ := range b {
			if board[i][j] == 1 {
				if countSuroundOnes(board, i, j) < 2 {
					dc[i][j] = 0
				} else if countSuroundOnes(board, i, j) <= 3 {
					dc[i][j] = board[i][j]
				} else {
					dc[i][j] = 0
				}
			}

			if board[i][j] == 0 {
				if countSuroundOnes(board, i, j) == 3 {
					dc[i][j] = 1
				}
			}
		}
	}

	for i, line := range dc {
		board[i] = line
	}
}

func main() {
	input := [][]int{
		{0,1,0},
		{0,0,1},
		{1,1,1},
		{0,0,0}}

	gameOfLife(input)

	fmt.Println(input)

}