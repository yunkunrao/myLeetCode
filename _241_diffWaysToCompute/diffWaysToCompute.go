package main

import "fmt"

type Operation struct {
	Pos    int
	Symbol rune
}

func diffWaysToCompute(input string) []int {
	var results = make([]int, 0)
	var positions []Operation
	var nums []int
	var num int

	for pos, c := range input {
		if c >= '0' && c <= '9' {
			num *= 10
			num += int(c - '0')
		}else {
			positions = append(positions, Operation{
				Pos:    pos,
				Symbol: c,
			})
			nums = append(nums, num)
			num = 0
		}
	}
	nums = append(nums, num)
	if len(nums) == 1 {
		results = append(results, nums[0])
		return results
	}

	for _, pos := range positions {
		for _, leftVal := range diffWaysToCompute(input[:pos.Pos]) {
			for _, rightVal := range diffWaysToCompute(input[pos.Pos+1:]) {
				if input[pos.Pos] == '+' {
					results = append(results, leftVal+rightVal)
				} else if input[pos.Pos] == '-' {
					results = append(results, leftVal-rightVal)
				} else if input[pos.Pos] == '*' {
					results = append(results, leftVal*rightVal)
				} else if input[pos.Pos] == '/' {
					results = append(results, leftVal/rightVal)
				} else {
					panic("invalid operator")
				}
			}
		}
	}

	return results
}

func main() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
