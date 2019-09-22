package _797_allPathsSourceTarget

func allPaths(from, to int, graph[][]int)[][]int {
	var paths [][]int

	if from == to {
		paths = append(paths, []int{from})
		return paths
	}


	for _, start := range graph[from] {
		for _, path := range allPaths(start, to, graph) {
			paths = append(paths, append([]int{from}, path...))
		}
	}

	return paths
}

func allPathsSourceTarget(graph [][]int) [][]int {

	var total = len(graph)

	if total == 0 {
		return nil
	}

	return allPaths(0, total-1, graph)

}
