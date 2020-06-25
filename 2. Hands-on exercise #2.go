package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeme(p *person) {
	p.first = "Miss"
	//*(p).last = "MoneyPenny"    this *(p).last and this p.last are same
	//p.age = 28
	fmt.Println(*p)
}
func main() {
	p := person{
		"James",
		"Bond",
		34,
	}
	fmt.Println(p)
	changeme(&p)
	fmt.Println(p)
}
