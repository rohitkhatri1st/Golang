package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Printf("%d\n%b\n%#x", a, a, a)	//%b and %#x are the format specifiers of binary and hex value starting with 0x respectively.
}	//%x is the format specifier for hexadecimal value without 0x.
