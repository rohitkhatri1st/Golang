package main

import (
	"fmt"
)

func main() {
	switch {
	case (2 == 3):
		fmt.Println("ASF")
	case (2 == 2):
		fmt.Println("ABC")
		fallthrough
	case (3 == 4):
		fmt.Println("Not")
	case (4 == 4):
		fmt.Println("Yes")
	}
}
