// 2. Grade Nilai

package main

import "fmt"

func main() {
	nilai := 5

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai >= 65 && nilai <= 79 {
		fmt.Println("B")
	} else if nilai >= 50 && nilai <= 64 {
		fmt.Println("C")
	} else if nilai >= 35 && nilai <= 49 {
		fmt.Println("D")
	} else if nilai >= 0 && nilai <= 34 {
		fmt.Println("E")
	} else if nilai <= 0 || nilai >= 100 {
		fmt.Println("Nilai Invalid")
	}
}
