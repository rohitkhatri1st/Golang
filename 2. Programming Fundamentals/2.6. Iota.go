package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 2016	// Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.
		b	//It will store value 2017 since iota is been used to assign value to a.
		c
		d
	)
	fmt.Println(a, b, c, d)
}

// output:
// 2016 2017 2018 2019
//
