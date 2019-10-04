package _1104_pathInZigZagTree

func pathInZigZagTree(label int) []int {
	var path []int
	var helper = []int{2}
	var count int = 0
	var tmp = label

	for tmp > 0 {
		tmp /= 2
		if tmp == 0 {
			break
		}
		count++
	}

	for i:=1; i<=count; i++ {
		next := 2*helper[i-1]+1
		helper = append(helper, next)
	}

	for {
		path = append([]int{label}, path...)
		label /= 2
		if label == 0 {
			break
		}
		label = helper[count-1] - label
		count--
	}
	return path
}