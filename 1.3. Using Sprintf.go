package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	a := fmt.Sprintf("%v %v %v", x, y, z)	//Sprintf formats according to a format specifier and RETURNS the resulting STRING.
	//%v denotes default format of the value
	fmt.Println(a)
}
