package main

import "fmt"

// L   | C   |
// E T | O E |
// E   | D   |
func convert(s string, numRows int) string {

	if len(s) < 2 || numRows < 2{
		return s
	}

	ret := ""
	singleVsize := 2*numRows - 2
	rowBeginOffset := 0
	rowEndOffset := singleVsize - rowBeginOffset
	vNums := len(s) / singleVsize
	rem := len(s) % singleVsize

	for {
		for i:= 0; i < vNums; i++ {
			singleV := s[i*singleVsize:(i+1)*singleVsize]

			ret += string(singleV[rowBeginOffset])

			if rowBeginOffset == rowEndOffset {
				continue
			}

			if rowBeginOffset != 0 {
				ret += string(singleV[rowEndOffset])
			}
		}

		left := s[len(s)-rem:]



		if rowBeginOffset < len(left) {
			ret += string(left[rowBeginOffset])
		}

		if rowBeginOffset != rowEndOffset && rowEndOffset < len(left) {
			ret += string(left[rowEndOffset])
		}

		rowBeginOffset++
		rowEndOffset--

		if rowBeginOffset > rowEndOffset {
			break
		}
	}
	return ret
}

func main() {
	var succeed bool = true

	testData := make(map[string]map[int]string)
	testData["LEETCODEISHIRING"] = make(map[int]string)
	testData["A"] = make(map[int]string)
	testData["AB"] = make(map[int]string)

	testData["LEETCODEISHIRING"][3] = "LCIRETOESIIGEDHN"
	testData["LEETCODEISHIRING"][4] = "LDREOEIIECIHNTSG"
	testData["A"][1] = "A"
	testData["AB"][1] = "AB"

	for s := range testData {
		for numRows := range testData[s]{
			if convert(s, numRows) != testData[s][numRows] {
				succeed = false
			}
		}
	}

	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

}
