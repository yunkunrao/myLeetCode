package _96_numTrees

func numTrees(n int) int {

	var array = make([]int, n+1)

	for i:=0; i<=n; i++ {
		if i < 2 {
			array[i] = 1
			continue
		}
		for j:=0; j<i; j++ {
			array[i] += array[j] * array[i-1-j]
		}
	}

	return array[n]
}