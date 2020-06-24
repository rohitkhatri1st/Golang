package main

import "fmt"

func main() {
	if 42 < 54 {//syntax of if is (if  var declaration;  condition). The expression is optional.
		fmt.Println("42 is less")
	} else {//else needs to be written right after end parentheses of if.
		fmt.Println("54 is less")
	}
}
