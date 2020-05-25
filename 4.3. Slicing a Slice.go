package main

import (
	"fmt"
)

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range s {
		fmt.Println(i, v)

	}
	fmt.Println(s[:5]) 	//slicing s from start to element 4
	fmt.Println(s[5:])	//slicing s from  element 5 to end.
	fmt.Println(s[2:7])	//slicing s from  element 2(element to included) to element 6(element 7 excluded)
	fmt.Println(s[1:6])	//slicing s from  element 1 to element 6
}
