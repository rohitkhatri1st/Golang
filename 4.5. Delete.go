package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)//This deletes elements from index 3 to index 5.
	fmt.Println(x)//Ultimately we are appending x's elements from start to index number 2 and then from index number 6 to end.
}
