package main

import "fmt"

func main() {
	// A composit literal looks like: x:= type{values}
	// A slice allows you to group together values of the same type
	x := []int{4, 5, 6, 7, 8, 42}

	// we can append (add) numbers to a slice
	x = append(x, 43, 44, 56)
	fmt.Println(x)

	// we can append two slices?
	// ... means unfurl the slice
	y := []int{50, 55, 56, 57, 58}
	z := append(x, y...)
	fmt.Println(z)

}
