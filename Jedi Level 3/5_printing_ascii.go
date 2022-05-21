package main

import "fmt"

func main() {

	// loop from 33 to 122
	x := 65
	for {
		if x > 90 {
			break
		}
		y := string(x)
		fmt.Printf("x = %d \t in ascii %s\n", x, y)
		x++
	}
	fmt.Println("end")

}
