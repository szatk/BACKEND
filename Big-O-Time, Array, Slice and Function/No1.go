package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number < 2 {
		return false
	} else {
		sqrtn := int(math.Sqrt(float64(number)))
		for i := 2; i <= sqrtn; i++ {
			if int(number)%i == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007)) // true

	fmt.Println(primeNumber(1500450271)) // true

	fmt.Println(primeNumber(1000000000)) // false

	fmt.Println(primeNumber(10000000019)) // true

	fmt.Println(primeNumber(10000000033)) // true

	fmt.Println(primeNumber(2))

	fmt.Println(primeNumber(7))

	fmt.Println(primeNumber(10))

	fmt.Println(primeNumber(10))

	fmt.Println(primeNumber(10))
}
