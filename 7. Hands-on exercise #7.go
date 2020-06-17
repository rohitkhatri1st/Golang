package main

import "fmt"

func main() {
	i := func(x string) string {
		fmt.Print("Hello Everyone Myself, ", x, "...\n")
		return "I want you all to meet Miss MoneyPenny"
	}
	fmt.Println(i("James Bond"))
}
