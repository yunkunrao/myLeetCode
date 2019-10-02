package _89_grayCode

func grayCode(n int) []int {
	// 000
	// 001
	// 011
	// 010
	// 110
	// 111
	// 101
	// 100

	var ret []int
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	gList := grayCode(n-1)

	for i:=0; i<len(gList); i++ {
		ret = append(ret, gList[i])
	}
	for i:=len(gList)-1; i>=0; i-- {
		tmp := (1 << (uint(n)-1)) + gList[i]
		ret = append(ret, tmp)
	}
	return ret
}