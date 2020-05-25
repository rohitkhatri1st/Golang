package main

import (
	"fmt"
)
//rune codepoint is an alias for int32
func main() {
	for i := 1; i < 27; i++ {
		fmt.Printf("%d\n\t%#U\n\t%#U\n\t%#U\n", i, i+64, i+64, i+64)	//%#U prints rune code value of passed value.
	}
}
