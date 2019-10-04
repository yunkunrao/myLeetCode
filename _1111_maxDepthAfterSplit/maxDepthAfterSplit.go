package _1111_maxDepthAfterSplit

func maxDepthAfterSplit(seq string) []int {
	l, r := 0, 0
	ret := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			ret[i] = l
			l = (l + 1)%2
		} else {
			ret[i] = r
			r = (r + 1)%2
		}
	}
	return ret
}
