package main

import "fmt"

func main() {
	/*
		Creating a slice of a slice
	*/

	xp := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James"}}

	fmt.Println("\nITERATING OVER THE SLICE:")

	for i, v := range xp {
		fmt.Println(i, v)
	}

	fmt.Println("\nITERATING OVER THE RECORDS:")

	for i := 0; i < len(xp); i++ {
		fmt.Println("- RECORD ", i)
		for ii, v := range xp[i] {
			fmt.Println(ii, v)
		}
	}
}
