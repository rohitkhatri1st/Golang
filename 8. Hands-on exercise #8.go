package main

import (
	"fmt"
)

func main() {
	switch {
	case (2 == 3):
		fmt.Println("1")
	case ("as" == "asd"):
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}
}
