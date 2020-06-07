package main

import "fmt"

func main() {
	p1 := struct {
		name   []string
		age    int
		family map[string][]string
	}{
		name: []string{"James", "Bond"},
		age:  32,
		family: map[string][]string{
			"James": []string{"Miss", "Moneypenny"},
			"Bond":  []string{"No", "Family"},
		},
	}
	fmt.Println(p1.name, p1.age)
	for k, v := range p1.family {
		fmt.Printf("%v's family is %v\n", k, v)
	}
}
