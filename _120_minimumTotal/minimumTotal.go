package main

import "fmt"

// 从i,j出发,cur目前的路径和
// func helper(triangle [][]int, i,j int, dp[][]int) int {

//     if i == len(triangle)-1 {
//         return dp[i][j]
//     }

//     t1 := dp[i][j] + helper(triangle, i+1, j)
//     t2 := dp[i][j] + helper(triangle, i+1, j+1)

//     if t1 < t2 {
//         return t1
//     } else {
//         return t2
//     }
// }

func minimumTotal(triangle [][]int) int {
	N := len(triangle)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	for j := 0; j < N; j++ {
		dp[N-1][j] = triangle[N-1][j]
	}

	for i := N - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if dp[i+1][j] > dp[i+1][j+1] {
				dp[i][j] = dp[i+1][j+1] + triangle[i][j]
			} else {
				dp[i][j] = dp[i+1][j] + triangle[i][j]
			}
		}
	}

	return dp[0][0]
}

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	fmt.Println(minimumTotal(triangle))
}
