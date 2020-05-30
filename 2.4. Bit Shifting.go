package main

import (
	"fmt"
)
//bit shifting takes binary input and then shifts one bit.
func main() {
	a := 42	// 42=101010
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1	//Shifting to left 1 bit of a.
	fmt.Printf("%d\t%b\t%#x", b, b, b)	//b=1010100=84

}
