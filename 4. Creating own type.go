package main

import (
	"fmt"
)

type mynew int

var x mynew

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
