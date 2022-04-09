package main

import "fmt"

type dollar int

var x dollar

func main() {

	y := fmt.Sprintf("The value of x is %v", x)
	fmt.Print("\n", y, "\n")

	x = 42

	fmt.Print(x, "\n")

}
