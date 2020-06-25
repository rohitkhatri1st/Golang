package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Icecream", "Sunsets"},
	}
	m["Mcleod_Todd"] = []string{"Coding", "Teaching", "Movies"}
	for k, v := range m {
		fmt.Println("The record of ", k, ":")
		for i, j := range v {
			fmt.Println("\t", i, j)
		}
	}
	delete(m, "Mcleod_Todd")//Deletes Mcleod_Todd key and Corresponding Values from m.
	for k, v := range m {
		fmt.Println(k, v)
	}

}
