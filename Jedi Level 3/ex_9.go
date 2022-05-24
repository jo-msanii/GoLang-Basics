package main

import (
	"fmt"
)

func main() {

	// Create a program that uses a switch statement
	// with the switch expression specified as a variable
	// of TYPE string with the IDENTIFIER "favSport"

	var favSport string = "hockey"

	switch favSport {
	case "football":
		fmt.Printf("You're hooligan who like gentleman's games")
	case "rugby":
		fmt.Printf("You're gentleman who likes hooligan's games")
	case "Fencing":
		fmt.Printf("You're right posho ain't ya!")
	default:
		fmt.Printf("You favourite sport is %v. We haven't a word to say about you!", favSport)
	}

	fmt.Println("\nEnd")

}
