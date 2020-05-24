package main

import (
	"fmt"
)

func main() {
	arr := [5]int{42, 43, 44, 45, 46}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)

}
