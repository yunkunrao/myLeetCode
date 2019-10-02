package main

import (
	"fmt"
	"strconv"
)

type complexNum struct {
	Real int
	Img int
}

func str2complex(a string) complexNum {
	var num complexNum
	var realNeg bool = false
	var imgNeg bool = false
	var index int = 0
	aStr := []rune(a)

	if aStr[0] == '-' {
		realNeg = true
		aStr = aStr[1:]
	}

	for index < len(aStr) {
		if aStr[index] >= '0' && aStr[index] <= '9' {
			num.Real *= 10
			num.Real += int(aStr[index] - '0')
		} else {
			break
		}
		index++
	}

	aStr = aStr[index:]

	if aStr[1] != '-' {
		index = 1
		imgNeg = false
	} else {
		index = 2
		imgNeg = true
	}

	aStr = aStr[index:]

	index = 0
	for index < len(aStr) {
		if aStr[index] >= '0' && aStr[index] <= '9' {
			num.Img *= 10
			num.Img += int(aStr[index] - '0')
		} else {
			break
		}
		index++
	}

	if realNeg {
		num.Real *= -1
	}

	if imgNeg {
		num.Img *= -1
	}

	return num
}

func complex2Str(num complexNum) string {
	return strconv.Itoa(num.Real) + "+" + strconv.Itoa(num.Img) + "i"
}

func complexNumberMultiply(a string, b string) string {
	var num complexNum

	aa := str2complex(a)
	bb := str2complex(b)

	num.Real = aa.Real*bb.Real - aa.Img*bb.Img
	num.Img = aa.Real*bb.Img + aa.Img*bb.Real

	return complex2Str(num)
}

func main() {




	a := "78+-76i"
	b := "-86+72i"
	fmt.Println(complexNumberMultiply(a, b))
}