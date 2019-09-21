package _709_toLowerCase

func toLowerCase(str string) string {

	var ret []rune

	for _, c := range str {
		cStr := string(c)
		if cStr >= "A" && cStr <= "Z" {
			c = c - 'A' + 'a'
		}
		ret = append(ret, c)
	}

	return string(ret)
}