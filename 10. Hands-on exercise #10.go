package main

import "fmt"

var a string = " a's scope is full package"

func main() {
	x := "x's scope is only main Function"
	for i := "i's scope is only for loop"; ; {
		fmt.Println(i)
		break
	}
	//fmt.Println(i) //Gives Error
	{
		y := "y's scope is only between enclosing parentheses"
		fmt.Println(y)
	}
	fmt.Println(x)
	//fmt.Println(y) //Gives Error
	scopeofa()
}
func scopeofa() {
	fmt.Println(a)
	//fmt.Println(x) //Gives Error
}
