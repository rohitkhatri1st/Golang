package main

import "fmt"

func main() {
	x := "james"
	if 42 > 54 {
		fmt.Println("42 is less")	
	} else if x == "James" {//Works if it is just after the parentheses
		fmt.Println("I am James")

	} else {
		fmt.Println("54 is there.")
	}
}
