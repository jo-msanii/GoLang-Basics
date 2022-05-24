package main

import (
	"fmt"
)

func main() {

	// Create a program that uses a switch statement
	// with no switch expression specified

	v := "hope"
	switch {
	case v == "joy":
		fmt.Printf("You are feeling joy")
	case v == "peace":
		fmt.Printf("You are feeling peace")
	case v == "hope":
		fmt.Printf("You are feeling hope")
	case v == "anger":
		fmt.Printf("You are feeling anger")
	default:
		fmt.Printf("You are feeling %v", v)
	}

	fmt.Println("\nend")

}
