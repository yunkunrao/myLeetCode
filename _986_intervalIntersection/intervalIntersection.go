package main

import "fmt"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var ret [][]int
	var i, j = 0, 0

	na := len(A)
	nb := len(B)

	for i<na && j<nb {
		if A[i][1] <= B[j][1] {
			// A   1----3
			// B      2----4
			//
			if B[j][0] <= A[i][1] {
				if B[j][0] >= A[i][0] {
					ret = append(ret, []int{B[j][0], A[i][1]})
				} else {
					ret = append(ret, []int{A[i][0], A[i][1]})
				}

			} else {

			}
			i++
		} else {
			// A     2----4
			// B  1-----3
			//
			if A[i][0] <= B[j][1] {
				if A[i][0] >= B[j][0] {
					ret = append(ret, []int{A[i][0], B[j][1]})
				} else {
					ret = append(ret, []int{B[j][0], B[j][1]})
				}
			} else {

			}
			j++
		}

	}

	return ret
}

func main() {

	A := [][]int{
		{5, 9},
	}

	B := [][]int{
		{3, 10},
	}

	fmt.Println(intervalIntersection(A, B))
}