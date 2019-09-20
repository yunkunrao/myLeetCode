package main

import "fmt"

func defangIPaddr(address string) string {
	//parts := strings.Split(address, ".")
	//return strings.Join(parts, "[.]")
	var ret string
	var tmp []rune
	for _, c := range address {
		if string(c) != "." {
			tmp = append(tmp, c)
		} else {
			ret += string(tmp) + "[.]"
			tmp = nil
		}
	}
	if len(tmp) != 0 {
		ret += string(tmp)
	}

	return ret
}

func main() {
	testData := make(map[string]string)
	testData["1.1.1.1"] = "1[.]1[.]1[.]1"

	for k, v := range testData {
		if defangIPaddr(k) == v {
			fmt.Println("successed!")
		} else {
			fmt.Println("Failed!", defangIPaddr(k) , v)
		}
	}
}
