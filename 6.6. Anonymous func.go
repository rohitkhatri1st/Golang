package main

import "fmt"

func main() {
	func(x string, y int) {
		for ; y >= 0; y-- {
			if y%5 == 0 {
				fmt.Println("This is ", y, "times of ", x)
			}
		}

	}("Love", 200)
}
