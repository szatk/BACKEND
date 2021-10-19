package main

import "fmt"

func palindrome(input string) bool {
	for i := 0; i < len(input); i++ {
		j := len(input) - 1 - i
		if input[i] != input[j] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(palindrome("civic")) // true

	fmt.Println(palindrome("katak")) // true

	fmt.Println(palindrome("kasur rusak")) // true

	fmt.Println(palindrome("mistar")) // false

	fmt.Println(palindrome("lion")) // false

	fmt.Println("Ketikkan Suatu Kata atau Frasa")
	var kata string
	fmt.Scanln(&kata)

	if palindrome(kata) == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan palindrome")
	}
}
