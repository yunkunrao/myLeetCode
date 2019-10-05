package _791_customSortString

func customSortString(S string, T string) string {
	var ret []rune
	var stat = make(map[rune]int, 0)

	for _, s := range S {
		stat[s] = 0
	}

	for _, t := range T {
		if _, ok := stat[t]; ok {
			stat[t]++
		} else {
			ret = append(ret, t)
		}
	}

	for _, s := range S {
		if stat[s] >= 1 {
			for i:=0; i<stat[s]; i++ {
				ret = append(ret, s)
			}
		}
	}

	return string(ret)
}