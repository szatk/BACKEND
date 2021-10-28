package main

import "fmt"

var fiboCollect = []int{0, 1, 1}

func fibo(n int) int {

	// your code here
	var result int

	if len(fiboCollect) > n {
		return fiboCollect[n]
	} else {
		result = fibo(n-1) + fibo(n-2)
	}

	fiboCollect = append(fiboCollect, result)
	return result
}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 34

	fmt.Println(fibo(10)) // 55
	// fmt.Println(fiboCollect)

	fmt.Println(fibo(50)) // 55
}
