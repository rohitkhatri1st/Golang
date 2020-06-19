package main

import "fmt"

func main() {
	g := func(a int, b int) int {
		return a + b
	}

	fmt.Println(foo(g, 25, 45))
}
func foo(f func(a int, b int) int, a, b int) int {
	i := f(a, b)
	return (2 * i)
}
