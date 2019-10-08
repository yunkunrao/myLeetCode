package _969_pancakeSort

func flip(A []int, begin, end int) {
	for begin < end {
		tmp := A[begin]
		A[begin] = A[end]
		A[end] = tmp
		begin++
		end--
	}
}

func pancakeSort(A []int) []int {
	var ret []int
	if len(A) == 0 {
		return nil
	}

	var max int = -1
	var index int = -1
	for i, a := range A {
		if a > max {
			max = a
			index = i
		}
	}

	if index != 0 {
		flip(A, 0, index)
		ret = append(ret, index+1)
	}
	if len(A) != 1 {
		flip(A, 0, len(A)-1)
		ret = append(ret, len(A))
	}

	ret = append(ret, pancakeSort(A[:len(A)-1])...)
	return ret
}
