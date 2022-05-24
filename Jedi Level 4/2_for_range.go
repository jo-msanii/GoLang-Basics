package main

import "fmt"

func main() {
	// A composit literal looks like: x:= type{values}
	// A slice allows you to group together values of the same type
	x := []int{4, 5, 6, 7, 8, 42}

	// Get length of the array
	fmt.Println(len(x))

	// Printhte entire array
	fmt.Println(x)

	// get content of a specific location
	fmt.Println(x[3])

	// Loop over a composite literal
	for i, v := range x {
		fmt.Println(i, v)
	}
}
