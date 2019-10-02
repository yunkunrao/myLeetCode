package main

import "fmt"

func haveCommonRune(part1, part2 []rune) bool {
	var statMap = make(map[rune]bool)

	for _, p1 := range part1 {
		statMap[p1] = true
	}

	for _, p2 :=range part2 {
		if _, ok := statMap[p2]; ok {
			return true
		}
	}

	return false
}

func partitionLabels(S string) []int {
	var ret []int

	if len(S) == 0 {
		return nil
	}

	str := []rune(S)

	for i:=1; i < len(str); i++ {
		if haveCommonRune(str[:i], str[i:]) {
			continue
		} else {
			ret = append(ret, i)
			str = str[i:]
			ret = append(ret, partitionLabels(string(str[:]))...)
			str = nil
		}
	}
	if len(str) != 0 {
		ret = append(ret, len(str))
	}

	return ret
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
