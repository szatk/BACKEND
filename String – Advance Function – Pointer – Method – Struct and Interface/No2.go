package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	transform := func(r rune) rune {
		if offset > 26 {
			offset = offset % 26
		}

		s := int(r) + offset
		if s > 'z' {
			return rune(s - 26)
		} else if s < 'a' {
			return rune(s + 26)
		}
		return rune(s)
	}

	result := strings.Map(transform, input)

	return result
}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))

	// bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

	// mnopqrstuvwxyzabcdefghijkl

}
