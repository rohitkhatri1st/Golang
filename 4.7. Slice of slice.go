package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	xy := [][]string{x, y}//defining a slice of slice.
	fmt.Println(xy)
	for _, v := range xy {
		for j, k := range v {
			fmt.Println(j, k)
		}
	}
}
