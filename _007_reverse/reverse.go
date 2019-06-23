package main

import "fmt"

func reverse(x int) int {

	var neg bool
	var ret int

	max := int(^uint32(0) >> 1)

	tmp := make([]int,0)
	if x < 0 {
		neg = true
		x = -x
	}

	for {
		if x == 0 {
			break
		}

		tmp = append(tmp, x % 10)
		x /= 10
	}

	for _,val := range tmp {
		ret = ret * 10 + val
	}

	if ret > max {
		return 0
	}

	if neg {
		ret = -ret
	}

	return ret
}

func main() {

	var succeed bool = true

	testData := make(map[int]int)
	testData[123] = 321
	testData[-123] = -321
	testData[120] = 21

	for k,v := range testData {
		if v != reverse(k) {
			succeed = false
			break
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

}
