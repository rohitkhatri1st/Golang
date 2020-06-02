package main

import (
	"fmt"
)

func main() {
	for i := 0 /*INIT statement*/; i < 10001 /*CONDITION Expression*/; i++ /*POST Statement*/ {
		fmt.Println(i)
	}
}
//the INIT statement: executed before the first iteration
//the CONDITION expression: evaluated before every iteration
//the POST statement: executed at the end of every iteration
