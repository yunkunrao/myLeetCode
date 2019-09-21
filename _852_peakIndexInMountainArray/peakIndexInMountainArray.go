package _852_peakIndexInMountainArray

func peakIndexInMountainArray(A []int) int {

	for i, a := range A {
		if i == len(A) - 1 {
			return -1
		}

		if i > 0 && a > A[i-1] && a > A[i+1] {
			return i
		}
	}
	return -1
}
