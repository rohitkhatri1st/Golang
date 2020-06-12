package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(foo(a...))
	fmt.Println(bar(a))
}

func foo(x ...int) int {
	i := 0
	for _, v := range x {
		i += v
	}
	return i
}
func bar(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}
