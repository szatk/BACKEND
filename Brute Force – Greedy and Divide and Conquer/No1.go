package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// your code here
	for x := 1; x <= 10000; x++ {
		for y := 1; y <= 10000-x; y++ {
			for z := 1; z <= 10000/(x*y); z++ {
				if x+y+z == a {
					if x*y*z == b {
						if x*x+y*y+z*z == c {
							fmt.Println(x, y, z)
							return
						}
					}
				}
			}
		}
	}

	fmt.Println("no solution")
	return
}

// func SimpleEquations(x, y, z int) {
// 	for i := 1; i < 10000; i++ {
// 		if i+(i+1)+(i+2) == x && i*(i+1)*(i+2) == y && (i*i)+((i+1)*(i+1))+((i+2)*(i+2)) == z {
// 			fmt.Println(i, (i + 1), (i + 2))
// 			return
// 		}
// 	}
// 	fmt.Println("No Solution")
// }

// func SimpleEquations(a, b, c int) {

// 	// your code here
// 	var panjang int
// 	if c < a || c < b {
// 		fmt.Println("No Solution")
// 		return
// 	}
// 	if a > b {
// 		panjang = a
// 	} else {
// 		panjang = b
// 	}
// 	if panjang <= 5 {
// 		fmt.Println("No Solution")
// 		return
// 	}
// 	for i := 1; i <= (panjang/2)-2; i++ {
// 		if i+(i+1)+(i+2) == a && i*(i+1)*(i+2) == b && (i*i)+((i+1)*(i+1))+((i+2)*(i+2)) == c {
// 			fmt.Println(i, i+1, i+2)
// 			return
// 		}
// 		for j := i + 1; j <= (panjang/2)-1; j++ {
// 			if i+j+(j+1) == a && i*j*(j+1) == b && (i*i)+(j*j)+((j+1)*(j+1)) == c {
// 				fmt.Println(i, j, j+1)
// 				return
// 			}
// 			for k := j + 1; k <= (panjang / 2); k++ {
// 				if i+(j)+(k) == a && i*(j)*(k) == b && (i*i)+((j)*(j))+((k)*(k)) == c {
// 					fmt.Println(i, j, k)
// 					return
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println("No Solution")

// }

func main() {

	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3

	SimpleEquations(16, 50, 126) // no solution

}
