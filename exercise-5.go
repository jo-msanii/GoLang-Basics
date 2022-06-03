package main

import "fmt"

type dollar int

var x dollar
var y int

func main() {

	z := fmt.Sprintf("The value of x is %v", x)
	fmt.Print(z, "\n")

	x = 42

	fmt.Print(x, "\n")

	y = int(x)
	fmt.Println(y)

}
