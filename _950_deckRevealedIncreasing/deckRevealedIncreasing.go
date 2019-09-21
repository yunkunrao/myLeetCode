package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	clockwiseRotate := func(array []int) []int{
		if array == nil || len(array) < 2 {
			return array
		}
		tmp := array[len(array)-1]
		array = array[0:len(array)-1]
		array = append([]int{tmp}, array...)
		return array
	}

	var ret []int
	sort.Ints(deck)

	for i:=len(deck)-1; i >=0 ; i-- {
		ret = clockwiseRotate(ret)
		ret = append([]int{deck[i]}, ret...)
	}

	return ret
}

func main() {
	input := []int{17,13,11,7,5,3,2}
	output := []int{2,13,3,11,5,17,7}

	for i := range output {
		ret := deckRevealedIncreasing(input)
		if ret[i] != output[i] {
			fmt.Println("failed", deckRevealedIncreasing(input), "expected", output)
			return
		}
	}

	fmt.Println("successed")
}
