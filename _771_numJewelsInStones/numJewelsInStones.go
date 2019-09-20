package main

import "fmt"

// 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 
// S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
//J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

func numJewelsInStones(J string, S string) int {
	if len(J) == 0 || len(S) == 0 {
		return 0
	}

	cmap := make(map[string]int)

	for _, j := range J {
		cmap[string(j)] = 0
	}

	for _, s := range S {
		if _, ok := cmap[string(s)]; ok {
			cmap[string(s)]++
		}
	}

	var total int = 0
	for _, v := range cmap {
		total += v
	}

	return total
}

func main() {

	var succeed= true

	testData := make(map[string]map[string]int)
	testData["aA"] = make(map[string]int)
	testData["aA"]["aAAbbbb"] = 3
	testData["z"] = make(map[string]int)
	testData["z"]["ZZ"] = 0

	for k1, kv := range testData {
		for k2, v := range kv {
			if numJewelsInStones(k1, k2) != v {
				succeed = false
				break
			}
		}
	}

	if succeed {
		fmt.Println("successed!")
	} else {
		fmt.Println("failed")
	}
}