package main

import (
	"fmt"
)

var x int	//var has a package level scope while short declaration operator has main level scope.
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
