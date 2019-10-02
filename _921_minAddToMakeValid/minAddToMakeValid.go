package _921_minAddToMakeValid

func minAddToMakeValid(S string) int {
	var leftPos []int
	var rightPos []int
	var matchPos []int

	for i, s := range S {
		if s == '(' {
			leftPos = append(leftPos, i)
		}
		if s == ')' {
			rightPos = append(rightPos, i)
		}
	}

	var index = 0
	for _, l := range leftPos {
		for index < len(rightPos) {
			if l < rightPos[index] {
				matchPos = append(matchPos, l)
				matchPos = append(matchPos, rightPos[index])
				index++
				break
			}
			index++
		}
	}

	return len(S) - len(matchPos)
}