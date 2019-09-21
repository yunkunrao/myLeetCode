package _70_climbStairs

func climbStairs(n int) int {

	var stat = []int{0, 1, 2}

	for i:=3; i <= n; i++ {
		stat = append(stat, stat[i-2] + stat[i-1])
	}
	return stat[n]
}