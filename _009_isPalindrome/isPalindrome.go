package main

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {

		var str []rune

		for ; x > 0; x = x / 10 {
			var c rune
			c = rune('0' + x%10)
			str = append(str, c)
		}

		l := len(str)
		for i := 0; i < l/2; i++ {
			if str[i] != str[l-i-1] {
				return false
			}
		}

		return true

	}
}

func main() {

	var succeed = true

	testData := make(map[int]bool)
	testData[121] = true
	testData[-121] = false
	testData[10] = false

	for k, v := range testData {
		if v != isPalindrome(k) {
			succeed = false
			fmt.Println(k, v)
			break
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

}
