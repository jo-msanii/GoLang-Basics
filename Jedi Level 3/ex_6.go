package main

import (
	"fmt"
)

func main() {

	// Create a program that shows the if statement in action

	n := 0
	for {
		if n >= 10 {
			break
		} else if n > 5 {
			fmt.Printf("n is less than 5 at only %v\n", n)
		} else {
			fmt.Printf("n is greater than 5 at %v\n", n)
		}
		n++
	}

	fmt.Println("end")

}
