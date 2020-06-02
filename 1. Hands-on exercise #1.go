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
	fmt.Println(p1.first, p1.last, "Favourite icecreams are : ")
	for i, v := range p1.favice {
		fmt.Printf("\t %v %v\n", i, v)
	}
	fmt.Println(p2.first, p2.last, "Favourite icecreams are : ")
	for i, v := range p2.favice {
		fmt.Printf("\t %v %v\n", i, v)
	}
}
