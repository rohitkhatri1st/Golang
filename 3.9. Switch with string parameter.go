package main

import (
	"fmt"
)

func main() {
	favsport := "Cricket"
	switch favsport {
	case "Football":
		fmt.Println("1")
	case ("Cricket"):
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}
}
