package main

import (
	"fmt"
)

func main() {
	bd := 1998
	for {	//infinity loop until Returns or breaks
		fmt.Println(bd)
		bd++
		if bd > 2020 {
			break
		}
	}
}
