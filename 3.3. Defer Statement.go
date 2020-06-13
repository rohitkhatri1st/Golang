package main

import "fmt"

func main() {
	defer bar()
	foo()
}
func foo() {
	defer bar()
	fmt.Println("I am foo")

}
func bar() {
	fmt.Println("I am bar")
}
