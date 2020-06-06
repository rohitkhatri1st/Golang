package main

import (
	"fmt"
)

func main() {
	age := 22	//For loop in golang can be used with or without any conditions.
	for age > 0 { //There is no while loop in golang.
		fmt.Println(2021 - age)
		age--
	}
}
