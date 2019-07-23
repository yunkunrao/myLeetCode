package main

import "fmt"

// 给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

func intToRoman2(num int) string {

	var ret string

	if !(num >= 1 && num <= 3999) {
		return ""
	}

	num2c := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	n1000 := num / 1000
	for i := 0; i < n1000; i++ {
		ret += num2c[1000]
	}
	num %= 1000

	n900 := num / 900
	if n900 > 0 {
		ret += num2c[900]
	}
	num %= 900

	n500 := num / 500
	if n500 > 0 {
		ret += num2c[500]
	}
	num %= 500

	n400 := num / 400
	if n400 > 0 {
		ret += num2c[400]
	}
	num %= 400

	n100 := num / 100
	for i := 0; i < n100; i++ {
		ret += num2c[100]
	}
	num %= 100

	n90 := num / 90
	if n90 > 0 {
		ret += num2c[90]
	}
	num %= 90

	n50 := num / 50
	if n50 > 0 {
		ret += num2c[50]
	}
	num %= 50

	n40 := num / 40
	if n40 > 0 {
		ret += num2c[40]
	}
	num %= 40

	n10 := num / 10
	for i := 0; i < n10; i++ {
		ret += num2c[10]
	}
	num %= 10

	n9 := num / 9
	if n9 > 0 {
		ret += num2c[9]
	}
	num %= 9

	n5 := num / 5
	if n5 > 0 {
		ret += num2c[5]
	}
	num %= 5

	n4 := num / 4
	if n4 > 0 {
		ret += num2c[4]
	}
	num %= 4

	n1 := num / 1
	for i := 0; i < n1; i++ {
		ret += "I"
	}

	return ret
}

func intToRoman(num int) string {

	var ret string

	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	if !(num >= 1 && num <= 3999) {
		return ""
	}

	index := 0

	for num > 0 {
		if num-nums[index] >= 0 {
			num -= nums[index]
			ret += romans[index]
			continue
		}
		index++
	}
	return ret
}

func main() {

	var succeed = true

	testData := make(map[int]string)
	testData[3] = "III"
	testData[4] = "IV"
	testData[9] = "IX"
	testData[20] = "XX"
	testData[58] = "LVIII"
	testData[100] = "C"
	testData[1994] = "MCMXCIV"

	for k, v := range testData {
		if v != intToRoman(k) {
			succeed = false
			fmt.Println(k, v, intToRoman(k))
			break
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

}
