package _807_maxIncreaseKeepingSkyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var hMax = map[int]int{}
	var vMax = map[int]int{}
	var sum int = 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if hMax[i] < grid[i][j] {
				hMax[i] = grid[i][j]
			}
			if vMax[j] < grid[i][j] {
				vMax[j] = grid[i][j]
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			limit := hMax[i]
			if limit > vMax[j] {
				limit = vMax[j]
			}
			sum += limit - grid[i][j]
		}
	}

	return sum
}
