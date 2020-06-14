package main

import (
	"fmt"
)

func main() {
	switch/*No expression given Here means true*/ {
	case (2 == 3):
		fmt.Println("1")
	case ("as" == "asd"):
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}
}
