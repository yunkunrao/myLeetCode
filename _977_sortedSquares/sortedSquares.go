package main

import "fmt"

func sortedSquares(A []int) []int {
	var ret []int
	var tmp []int
	var flag bool = true
	var negIndex, notNegIndex int

	for i, a := range A {
		if (flag && A[0] >= 0 ) || (a >= 0 && A[i-1] < 0) {
			notNegIndex = i
			flag = false
		}

		tmp = append(tmp, a*a)
	}

	negIndex = notNegIndex - 1

	for ;negIndex >=0 && notNegIndex < len(tmp); {
		if tmp[negIndex] <= tmp[notNegIndex] {
			ret = append(ret, tmp[negIndex])
			negIndex--
		} else {
			ret = append(ret, tmp[notNegIndex])
			notNegIndex++
		}
	}

	for ;negIndex >= 0;negIndex-- {
		ret = append(ret, tmp[negIndex])
	}
	for ;notNegIndex < len(tmp); notNegIndex++ {
		ret = append(ret, tmp[notNegIndex])
	}

	return ret
}

func main() {

	testData := []int{0,1,3,4}

	fmt.Println(sortedSquares(testData))

}
