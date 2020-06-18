package main

import "fmt"

var chk = (48 > 24)

func main() {
	fmt.Printf("%T\n", foo())
	fmt.Println(foo()())

}
func foo() func() int {
	if chk == true {
		return func() int {
			fmt.Println("Hello")
			return 23
		}

	} else {
		return func() int {
			fmt.Println("So sad it's false")
			return 0
		}
	}
}

/*func add(a, b int) int {
	return (a + b)
}
func multiply(a, b int) int {
	return (a * b)
}*/
