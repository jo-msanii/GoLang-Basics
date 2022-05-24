package main

import "fmt"

func main() {

	// No Fall through
	switch {
	case false:
		fmt.Println("shouldn't print")
	case 2 == 4:
		fmt.Println("shouldn't print")
	case 4 == 4:
		fmt.Println("should print – no fallthrough – END CASE ONE")
	}

	// Fall through
	switch {
	case false:
		fmt.Println("shouldn't print")
	case 2 == 4:
		fmt.Println("shouldn't print")
	case 4 == 4:
		fmt.Println("should print")
		fallthrough
	case 5 == 5:
		fmt.Println("FALLTHROUGH PRINTED – END CASE TWO")
	}

	// Default
	switch {
	case false:
		fmt.Println("shouldn't print")
	case 2 == 4:
		fmt.Println("shouldn't print")
	case 3 == 4:
		fmt.Println("should print")
		fallthrough
	default:
		fmt.Println("Default – END CASE THREE")
	}

	// Switch on a value
	switch "job" {
	case "job":
		fmt.Println("Job")
	case "Jane":
		fmt.Println("Jane")
	case "Rose":
		fmt.Println("Rose")
	}

	// Switch on multiple values
	switch "job" {
	case "job", "ballard":
		fmt.Println("Job – multiple values switch")
	case "Jane":
		fmt.Println("Jane")
	case "Rose":
		fmt.Println("Rose")
	}
}
