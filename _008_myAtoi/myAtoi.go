package main

import (
	"fmt"
)

func myAtoi(str string) int {
	var neg, numStart bool
	var ret int

	INT_MAX := int(^uint32(0) >> 1)
	INT_MIN := ^INT_MAX

	for i := 0; i < len(str); i++ {
		var num int

		c := str[i]
		if !numStart && c == ' ' {
			continue
		}

		if !((c >= '0' && c <= '9') ||
			(!numStart && (c == '+' || c == '-'))) {
			break
		}

		if !numStart {
			if c == '-' {
				neg = true
				numStart = true
				continue
			} else if c == '+' {
				numStart = true
				continue
			} else if c >= '0' && c <= '9' {
				numStart = true
			}
		}

		num = int(c - '0')

		ret = 10*ret + num

		if ret > INT_MAX {
			if neg {
				return INT_MIN
			} else {
				return INT_MAX
			}
		}
	}

	if neg {
		ret = -ret
	}

	if ret > INT_MAX {
		return INT_MAX
	} else if ret < INT_MIN {
		return INT_MIN
	}

	return ret
}

func main() {

	var succeed bool = true

	testData := make(map[string]int)
	testData["42"] = 42
	testData["  -42"] = -42
	testData["4193 with words"] = 4193
	testData["-91283472332"] = -2147483648
	testData["   +0 123"] = 0
	testData["9223372036854775808"] = 2147483647
	testData["0-1"] = 0

	for k, v := range testData {
		if v != myAtoi(k) {
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
