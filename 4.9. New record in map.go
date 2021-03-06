package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Icecream", "Sunsets"},
	}
	m["Mcleod_Todd"] = []string{"Coding", "Teaching", "Movies"}//This is creating a new record with key Todd_Mcleod
	for k, v := range m {
		fmt.Println("The record of ", k, ":")
		for i, j := range v {
			fmt.Println("\t", i, j)
		}
	}
}
