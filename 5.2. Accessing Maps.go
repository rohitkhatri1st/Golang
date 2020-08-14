package main

import "fmt"

type person struct {
	first  string
	last   string
	favice []string
}

func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		favice: []string{"Mint", "Strawberry", "Chocolate"},
	}
	p2 := person{
		first:  "Todd",
		last:   "Mcleod",
		favice: []string{"Martini", "Bubblegum", "Vanilla"},
	}
	m := map[string]person{
		p1.last: p1, //creating map with key=last of p1 and value = p1
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last, "Favourite icecreams are")
		for i, j := range v.favice {
			fmt.Println("\t", i, j)

		}
		fmt.Println("______________________")
	}
}
