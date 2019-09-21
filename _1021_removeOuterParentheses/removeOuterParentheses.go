package _1021_removeOuterParentheses

func removeOuterParentheses(S string) string {

	var count int = 0
	var ret []rune

	for _, c := range S {

		if string(c) == "(" {
			count++
		} else if string(c) == ")" {
			count--
		}

		if (count == 1 && string(c) == "(") ||
			(count == 0 && string(c) == ")") {

		} else {
			ret = append(ret, c)
		}
	}

	return string(ret)
}
