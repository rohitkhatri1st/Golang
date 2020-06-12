package main

import "fmt"

func main() {
	x := "james"
	if 42 > 54 {
		fmt.Println("42 is less")	
	} else if x == "James" {//Works if it is just after the parentheses
		fmt.Println("I am James")
	//else if or else will not run if 'if' statement has executed. Only one of 'if','else if' or 'else' will execute.
	} else {//Works if it is just after the parentheses
		fmt.Println("54 is there.")
	}
}
