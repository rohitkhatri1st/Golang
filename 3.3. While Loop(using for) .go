package main

import (
	"fmt"
)

func main() {
	age := 22
	for age > 0 { //There is no while loop in golang.
		fmt.Println(2021 - age)
		age--
	}
}
