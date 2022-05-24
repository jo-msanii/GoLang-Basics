package main

import "fmt"

func main() {
	// A composit literal looks like: x:= type{values}
	// A slice allows you to group together values of the same type
	x := []int{4, 5, 6, 7, 8, 42}
	fmt.Printf("%T", x)
}
