package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james"/*This is key*/:      []string{"Shaken, not stirred", "Martinis", "Women"},//This is Value
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Icecream", "Sunsets"},
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
