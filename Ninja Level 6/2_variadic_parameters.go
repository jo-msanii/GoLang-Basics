package main

import "fmt"

// func (r receiever) identifier(parameter(s)) (return(s)){ ... }

func main() {
	total := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(total)

}

// Passed in an unlimited number of integars
// vairadic parameters are stored in slices of the type of the parameters submited
func foo(x ...int) int {

	// Print out the content of parameter x
	fmt.Println(x)

	// print ou the type of parameter x
	fmt.Printf("%T\n", x)

	// Let us create a function that adds up all the unlimited integars to yeild a result
	sum := 0
	for _, v := range x {
		sum += v
	}

	// return the resulting value
	return sum
}
