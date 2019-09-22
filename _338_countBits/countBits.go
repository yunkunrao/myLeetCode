package _338_countBits

func countBits(num int) []int {

	var ret = make([]int, num+1)

	for i := 0; i <= num; i++ {
		ret[i] = ret[i>>1] + (i & 1)
	}

	return ret
}
