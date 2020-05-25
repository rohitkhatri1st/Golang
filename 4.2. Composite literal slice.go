package main

import (
	"fmt"
)

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}	// int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} is a compsite literal
	for i, v := range s {
		fmt.Println(i, v)

	}
	fmt.Printf("%T", s)
}
