package main

import "fmt"

func main() {

	x := 42
	if x == 40 {
		fmt.Println("our value is 40")
	} else if x == 41 {
		fmt.Println("our value is 41")
	} else {
		fmt.Println("You don't know our value!")
	}
}
