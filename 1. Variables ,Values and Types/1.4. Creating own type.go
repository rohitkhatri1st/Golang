package main

import (
	"fmt"
)

type mynew int	//type is used to define a new type. This statement says mynew has an underlying type "int".

var x mynew	//x is of type mynew and mynew has an underlying type "int".

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
