package main

import "fmt"

func main() {
	/*
		The benefit of make is that you can save memory by
		preventing the system from creating slices as you
		will have premade slices already made.
	*/

	// create the slice
	x := make([]int, 10, 100)

	// add into locations on the slice
	x[0] = 10
	fmt.Println(x)

	// Of course you can append to to it but that would
	// defeat the memory saving purposes
}
