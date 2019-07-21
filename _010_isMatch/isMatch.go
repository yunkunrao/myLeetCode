package main

import "fmt"

//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
//说明:
//
//s 可能为空，且只包含从 a-z 的小写字母。
//p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

func isMatch(s string, p string) bool {

	sArr := []rune(s)
	pArr := []rune(p)

	if len(pArr) == 0 {
		return len(sArr) == 0
	}

	var tmp string
	if len(sArr) > 0 {
		tmp = string(sArr[1:])
	}

	firstMatch := len(sArr) != 0 && (sArr[0] == pArr[0] || pArr[0] == '.')

	if len(pArr) >= 2 && pArr[1] == '*' {
		return isMatch(string(sArr), string(pArr[2:])) ||
			firstMatch && isMatch(tmp, string(pArr))

	} else {
		return firstMatch && isMatch(tmp, string(pArr[1:]))

	}

}

func main() {
	var succeed = true

	testData := make(map[string]map[string]bool)

	testData["aa"] = make(map[string]bool)
	testData["ab"] = make(map[string]bool)
	testData["aab"] = make(map[string]bool)
	testData["mississippi"] = make(map[string]bool)
	testData["aaba"] = make(map[string]bool)
	testData["aaa"] = make(map[string]bool)
	testData["a"] = make(map[string]bool)
	testData["bbbba"] = make(map[string]bool)

	//testData["aa"]["a"] = false
	//testData["aa"]["a*"] = true
	//testData["ab"][".*"] = true
	//testData["aab"]["c*a*b"] = true
	//testData["mississippi"]["mis*is*p*."] = false
	//testData["aaba"]["ab*a*c*a"] = false
	//testData["ab"][".*c"] = false
	testData["aaa"]["aaaa"] = false
	//testData["aaa"]["a*a"] = true
	//testData["aaa"]["ab*a*c*a"] = true
	//testData["a"]["ab*"] = true
	//testData["bbbba"][".*a*a"] = true

loop:
	for k1, v := range testData {
		for k2 := range v {
			if isMatch(k1, k2) != testData[k1][k2] {
				succeed = false
				fmt.Println("\n\n", k1, k2)
				break loop
			}
		}

	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}
}
