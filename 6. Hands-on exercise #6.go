package main

import (
	"fmt"
)

func main() {
	const (
		a = iota + 2016
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}

// output:
// 2016 2017 2018 2019
//
