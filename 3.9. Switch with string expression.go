package main

import (
	"fmt"
)

func main() {
	favsport := "Cricket"
	switch favsport/*favsport is the expression*/ {
	case "Football":
		fmt.Println("1")
	case ("Cricket")://Cricket is the string expression
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}
}
