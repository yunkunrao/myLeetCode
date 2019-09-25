package main

import (
	"fmt"
	"math"
)

func getAllPrimeFactors(num int) {
	var primeFactors []int
	var i int = 2
	for ; i <= num; i++ {
		if num % i == 0 {
			var flag bool = true
			for _, k := range primeFactors {
				if i % k == 0 && i != k {
					flag = false
					break
				}
			}
			if flag {
				primeFactors = append(primeFactors, i)
				num = num / i
				i = 1
			}
		}
	}

	for _, k := range primeFactors {
		fmt.Printf("%d ", k)
	}
}

func main() {
	//var i int
	//fmt.Scan(&i)
	//getAllPrimeFactors(i)
	fmt.Println(math.Sqrt(15*16))
}