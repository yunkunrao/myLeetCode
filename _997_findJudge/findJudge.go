package _997_findJudge

func findJudge(N int, trust [][]int) int {

	var in = make(map[int]int)
	var out = make(map[int]int)

	if N == 1 {
		return 1
	}

	// 统计每个顶点入度和出度
	for _, t := range trust {
		in[t[1]]++
		out[t[0]]++
	}

	for k, v := range in {
		if v == N-1 {
			if out[k] == 0 {
				return k
			}
		}
	}

	return -1
}
