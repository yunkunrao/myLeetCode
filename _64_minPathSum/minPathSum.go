package main

import "fmt"

func helper(grid [][]int, m, n int) int {
	var ret int = -1
	if m > len(grid)-1 || n > len(grid[0])-1 {
		return int(^uint(0)>>1)
	}

	if m == len(grid)-1 && n == len(grid[0])-1 {
		return grid[m][n]
	}

	t1 := helper(grid, m, n+1)
	if t1 != int(^uint(0)>>1) {
		t1 += grid[m][n]
	}

	t2 := helper(grid, m+1, n)
	if t2 != int(^uint(0)>>1) {
		t2 += grid[m][n]
	}

	if t1 < t2 {
		ret = t1
	} else {
		ret = t2
	}
	return ret
}

func minPathSum(grid [][]int) int {
	m := len(grid)-1
	n := len(grid[0])-1
	dp := make([][]int, len(grid))
	for i,_ := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	dp[m][n] = grid[m][n]
	for i:=m-1; i>=0; i-- {
		dp[i][n] = grid[i][n] + dp[i+1][n]
	}

	for j:=n-1; j>=0; j-- {
		dp[m][j] = grid[m][j] + dp[m][j+1]
	}

	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			t1 := grid[i][j] + dp[i][j+1]
			t2 := grid[i][j] + dp[i+1][j]
			if t1 > t2 {
				dp[i][j] = t2
			} else {
				dp[i][j] = t1
			}
		}
	}

	return dp[0][0]
}

func main() {

	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}
