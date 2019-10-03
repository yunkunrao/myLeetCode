package main

import (
	"fmt"
)

func stoneGame(piles []int) bool {
	N := len(piles)
	fDp := make([][]int, N)
	for i, _ := range fDp {
		fDp[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		fDp[i][i] = piles[i]
	}

	sDp := make([][]int, N)
	for j, _ := range sDp {
		sDp[j] = make([]int, N)
	}

	for d := 1; d < N; d++ {
		for i := 0; i < N; i++ {
			j := i + d
			if j >= N {
				continue
			}

			// update fdp
			fDp[i][j] = sDp[i+1][j] + piles[i]
			tmp := sDp[i][j-1] + piles[j]
			if fDp[i][j] < tmp {
				fDp[i][j] = tmp
			}

			// update sdp
			sDp[i][j] = fDp[i+1][j]
			tmp = fDp[i][j-1]
			if sDp[i][j] > tmp {
				sDp[i][j] = tmp
			}

		}
	}

	return fDp[0][N-1] > sDp[0][N-1]
}

func main() {
	input := []int{7, 7, 12, 16, 41, 48, 41, 48, 7, 12, 11, 39, 31, 8, 23, 11, 47, 25, 15, 23, 4, 17, 11, 50, 16, 50, 38, 34, 48, 27, 16, 24, 22, 48, 50, 10, 26, 27, 9, 43, 13, 42, 46, 24}
	fmt.Println(stoneGame(input))
}
