package _451_frequencySort

import "sort"

type CharInfo struct {
	C rune
	Count int
}

func frequencySort(s string) string {
	var stat = make(map[rune]*CharInfo, 0)
	var list []*CharInfo
	var ret []rune

	for _, ss := range s {
		if _, ok := stat[ss]; !ok {
			stat[ss] = &CharInfo{
				C: ss,
				Count: 1,
			}
		} else {
			stat[ss].Count++
		}
	}

	for _, v := range stat {
		list = append(list, v)
	}

	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Count > list[j].Count
	})

	for _, c := range list {
		for i:=0; i<c.Count; i++ {
			ret = append(ret, c.C)
		}
	}

	return string(ret)
}