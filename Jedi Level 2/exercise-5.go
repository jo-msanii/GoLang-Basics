package main

import "fmt"

func main() {

	a := `42\t43`
	fmt.Printf(a)

	b := "\n42\t43"
	fmt.Printf(b)
}
