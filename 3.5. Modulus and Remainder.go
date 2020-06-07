package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Dividing %v with 4 gives remainder %v\n", i, i%4)	//"%", is the modulus operator. It gives the remainder
										// when i is divided by 4.
	}
}
