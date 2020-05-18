package main

import (
	"fmt"
)

func main() {
	const (
		a     = 42     //Typed Constants
		b     = 43.6
		c     = "Khatri"
		d int = 23	//Untyped Constants
		e int = 45
		f int = 25
	)
	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("%T\t%T\t%t\t%t\t%t\t%t", a, b, c, d, e, f)

}
