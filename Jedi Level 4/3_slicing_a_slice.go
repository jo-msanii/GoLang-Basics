package main

import "fmt"

func main() {
	// A composit literal looks like: x:= type{values}
	// A slice allows you to group together values of the same type
	x := []int{4, 5, 6, 7, 8, 42}

	// We can access a slice by going to the position we want
	fmt.Println(x[3])

	// Colon operater
	// Start from the stated index
	fmt.Println(x[1:])
	// Start from the stated index and stop before the next stated index
	fmt.Println(x[1:3])

	// Fun exercise: Create a for loop but don't use the range
	for c := 0; c < len(x); c++ {
		fmt.Println(x[c])
	}

}
