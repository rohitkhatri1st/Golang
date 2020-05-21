package main

import (
	"fmt"
)

func main() {
	bd := 1998
	for {
		fmt.Println(bd)
		bd++
		if bd > 2020 {
			break
		}
	}
}
