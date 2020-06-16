package main

import (
	"fmt"
)

func main() {//Elements in parentheses seperated by comma (combined), are composite literals.
	arr := [5]int{42, 43, 44, 45, 46}	//int {42, 43, 44, 45, 46} is a composite literal and [5] defines length of array
	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)

}
