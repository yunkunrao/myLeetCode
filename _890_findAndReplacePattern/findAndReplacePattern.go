package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {
	var ret []string
	var length = len(words)
	var wordMap = make(map[rune]rune, 0)
	var wordReverseMap = make(map[rune]rune, 0)

	if length == 0 {
		return nil
	}
	if length == 1 {
		word := words[0]
		if len(word) != len(pattern) {
			return nil
		}

		for i, w := range word {
			val, ok := wordMap[w]
			_, ok2 := wordReverseMap[[]rune(pattern)[i]]
			if !ok && !ok2 {
				wordMap[w] = []rune(pattern)[i]
				wordReverseMap[[]rune(pattern)[i]] = w
			} else {
				if val == []rune(pattern)[i] {
					continue
				} else {
					return nil
				}
			}
		}
		ret = append(ret, word)
		return ret
	}

	ret = append(ret, findAndReplacePattern(words[:length/2], pattern)...)
	ret = append(ret, findAndReplacePattern(words[length/2:], pattern)...)
	return ret
}

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Println(findAndReplacePattern(words, pattern))
}
