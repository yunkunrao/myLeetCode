package main

import "fmt"

func helper(grid [][]int, visited [][]bool, i, j int) int {
	var ret int = -1
	m := len(grid)
	n := len(grid[0])

	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == 0 {
		return 0
	}

	di := []int{-1, 1, 0, 0}
	dj := []int{0, 0, -1, 1}

	visited[i][j] = true
	for k := 0; k < 4; k++ {
		tmp := helper(grid, visited, i+di[k], j+dj[k])
		if ret < tmp {
			ret = tmp
		}
	}
	visited[i][j] = false

	return ret + grid[i][j]
}

func getMaximumGold(grid [][]int) int {
	var ret int = -1
	var visited [][]bool

	m := len(grid)
	n := len(grid[0])

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp := helper(grid, visited, i, j)
			if tmp > ret {
				ret = tmp
			}
		}
	}

	return ret
}

func main() {

	grid := [][]int{
		{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20},
	}

	fmt.Println(getMaximumGold(grid))
}
