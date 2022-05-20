package main

import "fmt"

func main() {

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n\n", y, y)

	// If we are to increment
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\t\t\t\t%x\n", kb, kb, kb)
	fmt.Printf("%d\t\t\t%b\t\t\t%x\n", mb, mb, kb)
	fmt.Printf("%d\t\t%b\t\t%x\n", gb, gb, kb)

}
