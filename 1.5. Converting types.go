package main

import (
	"fmt"
)

type mynew int

var x mynew
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)	//converting x of type mynew to int and storing it in y
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
